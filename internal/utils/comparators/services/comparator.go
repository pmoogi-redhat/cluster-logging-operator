package services

import (
	"github.com/ViaQ/logerr/log"
	"github.com/openshift/cluster-logging-operator/internal/utils"
	v1 "k8s.io/api/core/v1"
	"reflect"
)

//AreSame compares for equality and return true equal otherwise false
func AreSame(current *v1.Service, desired *v1.Service) bool {
	log.V(3).Info("Comparing Services current to desired", "current", current, "desired", desired)

	if !utils.AreMapsSame(current.ObjectMeta.Labels, desired.ObjectMeta.Labels) {
		log.V(3).Info("Service label change", "current name", current.Name)
		return false
	}
	if !utils.AreMapsSame(current.Spec.Selector, desired.Spec.Selector) {
		log.V(3).Info("Service Selector change", "current name", current.Name)
		return false
	}
	if len(current.Spec.Ports) != len(desired.Spec.Ports) {
		return false
	}
	for i, port := range current.Spec.Ports {
		dPort := desired.Spec.Ports[i]
		if !reflect.DeepEqual(port, dPort) {
			return false
		}
	}

	return true
}
