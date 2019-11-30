package grafana

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeService(c *pulsarv1alpha1.PulsarCluster) *v1.Service {
	var serviceType v1.ServiceType
	if c.Spec.Monitor.Grafana.NodePort == 0 {
		serviceType = v1.ServiceTypeClusterIP
	} else {
		serviceType = v1.ServiceTypeNodePort
	}
	return &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      MakeServiceName(c),
			Namespace: c.Namespace,
			Labels:    pulsarv1alpha1.MakeAllLabels(c, pulsarv1alpha1.MonitorComponent, pulsarv1alpha1.MonitorGrafanaComponent),
		},
		Spec: v1.ServiceSpec{
			Ports:    makeServicePorts(c),
			Type:     serviceType,
			Selector: pulsarv1alpha1.MakeAllLabels(c, pulsarv1alpha1.MonitorComponent, pulsarv1alpha1.MonitorGrafanaComponent),
		},
	}
}

func MakeServiceName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-monitor-grafana-service", c.GetName())
}

func makeServicePorts(c *pulsarv1alpha1.PulsarCluster) []v1.ServicePort {
	if c.Spec.Monitor.Grafana.NodePort == 0 {
		return []v1.ServicePort{
			{
				Name: "grafana",
				Port: pulsarv1alpha1.PulsarGrafanaServerPort,
			},
		}
	} else {
		return []v1.ServicePort{
			{
				Name:     "grafana",
				NodePort: c.Spec.Monitor.Grafana.NodePort,
				Port:     pulsarv1alpha1.PulsarGrafanaServerPort,
			},
		}
	}
}
