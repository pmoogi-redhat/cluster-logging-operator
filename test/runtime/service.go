package runtime

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ServiceBuilder struct {
	Service *corev1.Service
}

func NewServiceBuilder(svc *corev1.Service) *ServiceBuilder {
	return &ServiceBuilder{
		Service: svc,
	}
}

func (builder *ServiceBuilder) WithSelector(selector map[string]string) *ServiceBuilder {
	builder.Service.Spec.Selector = selector
	return builder
}

func (builder *ServiceBuilder) AddServicePort(port int32, targetPort int) *ServiceBuilder {
	builder.Service.Spec.Ports = append(builder.Service.Spec.Ports, corev1.ServicePort{
		Port:       port,
		TargetPort: intstr.FromInt(targetPort),
	})
	return builder
}
