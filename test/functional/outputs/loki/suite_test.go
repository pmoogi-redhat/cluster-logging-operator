package loki

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFunctionalOutputs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ClusterLogging Functional Loki Output Suite")
}
