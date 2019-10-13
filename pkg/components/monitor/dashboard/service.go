package dashboard

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeService(c *pulsarv1alpha1.PulsarCluster) *v1.Service {
	return &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      MakeServiceName(c),
			Namespace: c.Namespace,
			Labels:    pulsarv1alpha1.MakeAllLabels(c, pulsarv1alpha1.MonitorComponent, pulsarv1alpha1.MonitorDashboardComponent),
		},
		Spec: v1.ServiceSpec{
			Ports:    makeServicePorts(c),
			Type:     v1.ServiceTypeNodePort,
			Selector: pulsarv1alpha1.MakeAllLabels(c, pulsarv1alpha1.MonitorComponent, pulsarv1alpha1.MonitorDashboardComponent),
		},
	}
}

func MakeServiceName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-monitor-dashboard-service", c.GetName())
}

func makeServicePorts(c *pulsarv1alpha1.PulsarCluster) []v1.ServicePort {
	return []v1.ServicePort{
		{
			Name:     "dashboard",
			NodePort: c.Spec.Monitor.DashboardPort,
			Port:     PulsarDashboardContainerPort,
		},
	}
}
