package fluent_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo"
	loggingv1 "github.com/openshift/cluster-logging-operator/pkg/apis/logging/v1"
	"github.com/openshift/cluster-logging-operator/test"
	"github.com/openshift/cluster-logging-operator/test/client"
	"github.com/openshift/cluster-logging-operator/test/helpers/cmd"
	"github.com/openshift/cluster-logging-operator/test/helpers/fluentd"
	. "github.com/openshift/cluster-logging-operator/test/matchers"
	"github.com/openshift/cluster-logging-operator/test/runtime"
	corev1 "k8s.io/api/core/v1"
)

const message = "My life is my message"

var (
	// Create/delete clusterlogging instance around entire suite.
	cl = runtime.NewClusterLogging()
	_  = BeforeSuite(func() { ExpectOK(client.Get().Recreate(cl)) })
	_  = AfterSuite(func() { _ = client.Get().Remove(cl) })
)

// Extend client.Test for tests using logGenerator and receiver.
type Test struct {
	*client.Test
	receiver     *fluentd.Receiver
	logGenerator *corev1.Pod
	group        test.FailGroup // Run cluster operations in parallel
}

func (t *Test) ReaderForSource(name string) *cmd.Reader {
	return t.receiver.Sources[name].TailReader()
}

func NewTest(message string) *Test {
	ct := client.NewTest()
	t := &Test{
		Test:         ct,
		receiver:     fluentd.NewReceiver(ct.NS.Name, "receiver"),
		logGenerator: runtime.NewLogGenerator(ct.NS.Name, "log-generator", 10000, 0, message),
	}
	t.group.Go(func() { ExpectOK(t.Create(t.logGenerator)) })
	return t
}

func (t *Test) Close() {
	defer t.Test.Close()
	t.group.Wait()
}

var _ = Describe("[ClusterLogForwarder]", func() {
	var t *Test
	BeforeEach(func() { t = NewTest(message) })
	AfterEach(func() { t.Close() })

	Context("with app/infra/audit receiver", func() {
		BeforeEach(func() {
			t.receiver.AddSource(&fluentd.Source{Name: "application", Type: "forward", Port: 24224})
			t.receiver.AddSource(&fluentd.Source{Name: "infrastructure", Type: "forward", Port: 24225})
			t.receiver.AddSource(&fluentd.Source{Name: "audit", Type: "forward", Port: 24226})
			t.group.Go(func() { ExpectOK(t.receiver.Create(t.Client)) })
		})

		It("forwards application logs only", func() {
			clf := runtime.NewClusterLogForwarder()
			addPipeline(clf, t.receiver.Sources["application"])
			t.group.Go(func() { ExpectOK(t.Recreate(clf)) })
			// Use ctx to let empty checks run until we get the expected lines.
			ctx, cancel := context.WithCancel(context.Background())
			t.group.Go(func() {
				defer cancel() // Cancel goroutines verifying empty when we get our logs.
				r := t.ReaderForSource("application")
				ExpectOK(r.ExpectLines(10, message, `{"viaq_index_name":"(inf|aud)`))
			})
			t.group.Go(func() { ExpectOK(t.ReaderForSource("infrastructure").ExpectEmpty(ctx)) })
			t.group.Go(func() { ExpectOK(t.ReaderForSource("audit").ExpectEmpty(ctx)) })
		})

		It("forwards infrastructure logs only", func() {
			clf := runtime.NewClusterLogForwarder()
			addPipeline(clf, t.receiver.Sources["infrastructure"])
			t.group.Go(func() { ExpectOK(t.Recreate(clf)) })
			// Use ctx to let empty checks run until we get the expected lines.
			ctx, cancel := context.WithCancel(context.Background())
			t.group.Go(func() {
				defer cancel()
				r := t.ReaderForSource("infrastructure")
				ExpectOK(r.ExpectLines(10, "", `{"viaq_index_name":"(app|aud)`))
			})
			t.group.Go(func() { ExpectOK(t.ReaderForSource("application").ExpectEmpty(ctx)) })
			t.group.Go(func() { ExpectOK(t.ReaderForSource("audit").ExpectEmpty(ctx)) })
		})

		It("forwards audit logs only", func() {
			clf := runtime.NewClusterLogForwarder()
			addPipeline(clf, t.receiver.Sources["audit"])
			// Do everything in parallel - by eventual consistency the test will pass.
			t.group.Go(func() { ExpectOK(t.Recreate(clf)) })
			// Use ctx to let empty checks run until we get the expected lines.
			ctx, cancel := context.WithCancel(context.Background())
			t.group.Go(func() {
				defer cancel()
				r := t.ReaderForSource("audit")
				ExpectOK(r.ExpectLines(10, "", `{"viaq_index_name":"(inf|app)`))
			})
			t.group.Go(func() { ExpectOK(t.ReaderForSource("application").ExpectEmpty(ctx)) })
			t.group.Go(func() { ExpectOK(t.ReaderForSource("infrastructure").ExpectEmpty(ctx)) })
		})

		It("forwards different types to different outputs with labels", func() {
			clf := runtime.NewClusterLogForwarder()
			for _, name := range []string{"application", "infrastructure", "audit"} {
				s := t.receiver.Sources[name]
				clf.Spec.Outputs = append(clf.Spec.Outputs, loggingv1.OutputSpec{
					Name: s.Name,
					Type: "fluentdForward",
					URL:  fmt.Sprintf("tcp://%v:%v", s.Host(), s.Port),
				})
				clf.Spec.Pipelines = append(clf.Spec.Pipelines, loggingv1.PipelineSpec{
					InputRefs:  []string{s.Name},
					OutputRefs: []string{s.Name},
					Labels:     map[string]string{"log-type": s.Name},
				})
			}

			// Do everything in parallel - by eventual consistency the test will pass.
			t.group.Go(func() { ExpectOK(t.Recreate(clf)) })
			for _, name := range []string{"application", "infrastructure", "audit"} {
				name := name // Don't bind to range variable
				r := t.ReaderForSource(name)
				t.group.Go(func() {
					ExpectOK(r.ExpectLines(3, fmt.Sprintf(`"log-type":%q`, name), `"log-type":`))
				})
			}
		})
	})
})

func addPipeline(clf *loggingv1.ClusterLogForwarder, s *fluentd.Source) {
	clf.Spec.Outputs = append(clf.Spec.Outputs, loggingv1.OutputSpec{
		Name: s.Name,
		Type: "fluentdForward",
		URL:  fmt.Sprintf("tcp://%v:%v", s.Host(), s.Port),
	})
	clf.Spec.Pipelines = append(clf.Spec.Pipelines,
		loggingv1.PipelineSpec{
			InputRefs:  []string{s.Name},
			OutputRefs: []string{s.Name},
		})
}
