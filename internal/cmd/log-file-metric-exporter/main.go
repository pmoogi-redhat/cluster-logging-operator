package main

import (
	"flag"
	"github.com/ViaQ/logerr/log"
	"os"
	"net/http"
	"github.com/openshift/cluster-logging-operator/internal/pkg/symnotify"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)




var (
 verbosity int = 0
)



type FileWatcher struct {
        watcher *symnotify.Watcher
        metrics *prometheus.CounterVec
        sizes  map[string]float64
        added  map[string]bool
}


func (w *FileWatcher) Update(path string) error {
        counter, err := w.metrics.GetMetricWithLabelValues(path)
        if err != nil {
                return err
        }
        stat, err := os.Stat(path)
        if err != nil {
                return err
        }
        if stat.IsDir() {
                return nil // Ignore directories
        }
        lastSize, size := float64(w.sizes[path]), float64(stat.Size())
        w.sizes[path] = size
        var add float64
        if size > lastSize {
                // File has grown, add the difference to the counter.
                add = size - lastSize
        } else if size < lastSize {
                // File truncated, starting over. Add the size.
                add = size
        }
        log.V(2).Info("For logfile in...", "path",path,"lastsize",lastSize,"currentsize",size,"addedbytes",add)
        counter.Add(add)
        return nil
}

func (w *FileWatcher) Watch( ) {

        
	for {
		//All logfiles with containername are added to the watcher
		//write event for these logfiles are being watched
		//create event gets issued for all new logfiles appear under logfilepathname /var/log/containers/
		//For the cases new log files added, old files moved, old files deleted, you need to add/remove them from watcher as whole dir added to the watcher
		//For new log files added write event is not getting issued
		
                e, err := w.watcher.Event()
                if err != nil {
		log.Error(err, "Watcher.Event returning err")
		os.Exit(1)
		}

                log.V(2).Info("Events notified for...", "e.Name",e.Name,"Event",e.Op)
                w.Update(e.Name)
        }
}


func main() {
        var dir string
        var addr string

	//directory to be watched out where symlinks to all logs files are present e.g. /var/log/containers/
	//debug option true or false
	//listening port where this go-app push prometheus registered metrics for further collected or reading by end prometheus server
	flag.StringVar(&dir, "dir", "/var/log/containers/", "Directory containing log files")
	flag.IntVar(&verbosity, "verbosity", 0, "set verbosity level")
	flag.StringVar(&addr, "http", ":2112", "HTTP service address where metrics are exposed")
	flag.Parse()


	log.SetLogLevel(verbosity)

        log.V(2).Info("Watching out logfiles dir ...", "dir",dir,"http",addr)


	//Get new watcher
        w := &FileWatcher {
                metrics: prometheus.NewCounterVec(prometheus.CounterOpts{
                        Name: "log_logged_bytes_total",
                        Help: "Total number of bytes written to a single log file path, accounting for rotations",
                }, []string{"path"}),
               sizes: make(map[string]float64),
               added: make(map[string]bool),

        }
	
	prometheus.Register(w.metrics)
	defer prometheus.Unregister(w.metrics)

	symwatcher, err := symnotify.NewWatcher()
	w.watcher = symwatcher
	
	if err != nil {
		log.Error(err, "NewFileWatcher error")
		os.Exit(1)
	}
	defer w.watcher.Close()
	//Add dir to watcher
	 w.watcher.Add(dir)

	go w.Watch()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(addr, nil)
}
