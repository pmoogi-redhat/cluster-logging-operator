package telemetry

import (
	"github.com/openshift/cluster-logging-operator/version"
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

// placeholder for keeping clo info which will be used for clo metrics update
type TData struct {
	CLInfo              map[string]string
	CLOutputType        map[string]string
	CollectorErrorCount float64
	CLFInfo             map[string]string
	CLFInputType        map[string]string
	CLFOutputType       map[string]string
}

// "0" stands for managedStatus and healthStatus true and healthy
func NewTD() *TData {
	return &TData{
		CLInfo:              map[string]string{"version": version.Version, "managedStatus": "0", "healthStatus": "0"},
		CLOutputType:        map[string]string{"elasticsearch": "0"},
		CollectorErrorCount: 0,
		CLFInfo:             map[string]string{"healthStatus": "0", "pipelineInfo": "1"},
		CLFInputType:        map[string]string{"application": "1", "audit": "0", "infrastructure": "0"},
		CLFOutputType:       map[string]string{"default": "1", "elasticsearch": "0", "fluentdForward": "0", "syslog": "0", "kafka": "0", "loki": "0", "cloudwatch": "0"},
	}
}

var (
	Data = NewTD()

	mCLInfo = NewInfoVec(
		"log_logging_info",
		"Clo version managementState healthState specific metric",
		[]string{"version", "managedStatus", "healthStatus"},
	)
	mCollectorErrorCount = NewInfoVec(
		"log_collector_error_count_total",
		"log collector total number of error counts ",
		[]string{"version"},
	)
	mCLFInfo = NewInfoVec(
		"log_forwarder_pipeline_info",
		"Clf healthState and pipelineInfo specific metric",
		[]string{"healthStatus", "pipelineInfo"},
	)

	mCLFInputType = NewInfoVec(
		"log_forwarder_input_info",
		"Clf input type specific metric",
		[]string{"application", "audit", "infrastructure"},
	)

	mCLFOutputType = NewInfoVec(
		"log_forwarder_output_info",
		"Clf output type specific metric",
		[]string{"default", "elasticsearch", "fluentdForward", "syslog", "kafka", "loki", "cloudwatch"},
	)

	MetricList = []prometheus.Collector{
		mCLInfo,
		mCollectorErrorCount,
		mCLFInfo,
		mCLFInputType,
		mCLFOutputType,
	}
)

func RegisterMetrics() error {

	for _, metric := range MetricList {
		metrics.Registry.MustRegister(metric)
	}

	return nil
}

func UpdateMetrics() error {

	CLInfo := Data.CLInfo
	CollectorErrorCount := Data.CollectorErrorCount
	CLFInfo := Data.CLFInfo
	CLFInputType := Data.CLFInputType
	CLFOutputType := Data.CLFOutputType

	mCLInfo.With(prometheus.Labels{
		"version":       CLInfo["version"],
		"managedStatus": CLInfo["managedStatus"],
		"healthStatus":  CLInfo["healthStatus"]}).Set(1)

	mCollectorErrorCount.With(prometheus.Labels{
		"version": CLInfo["version"]}).Set(CollectorErrorCount)

	mCLFInfo.With(prometheus.Labels{
		"healthStatus": CLFInfo["healthStatus"],
		"pipelineInfo": CLFInfo["pipelineInfo"]}).Set(1)

	mCLFInputType.With(prometheus.Labels{
		"application":    CLFInputType["application"],
		"audit":          CLFInputType["audit"],
		"infrastructure": CLFInputType["infrastructure"]}).Set(1)

	mCLFOutputType.With(prometheus.Labels{
		"default":        CLFOutputType["default"],
		"elasticsearch":  CLFOutputType["elasticsearch"],
		"fluentdForward": CLFOutputType["fluentdForward"],
		"syslog":         CLFOutputType["syslog"],
		"kafka":          CLFOutputType["kafka"],
		"loki":           CLFOutputType["loki"],
		"cloudwatch":     CLFOutputType["cloudwatch"]}).Set(1)

	return nil
}

func NewInfoVec(metricname string, metrichelp string, labelNames []string) *prometheus.GaugeVec {

	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricname,
			Help: metrichelp,
		},
		labelNames,
	)
}
