package proxy

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/components/broker"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeConfigMap(c *pulsarv1alpha1.PulsarCluster) *v1.ConfigMap {
	brokerServiceName := broker.MakeServiceName(c)
	webServiceUrl := fmt.Sprintf("http://%s.%s.%s:8000", brokerServiceName, c.Namespace, pulsarv1alpha1.ServiceDomain)
	brokerServiceUrl := fmt.Sprintf("pulsar://%s.%s.%s:6650", brokerServiceName, c.Namespace, pulsarv1alpha1.ServiceDomain)

	return &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      MakeConfigMapName(c),
			Namespace: c.Namespace,
		},
		Data: map[string]string{
			"PULSAR_MEM":          PulsarMemData,
			"brokerServiceURL":    brokerServiceUrl,
			"brokerWebServiceURL": webServiceUrl,
			"clusterName":         c.GetName(),
		},
	}
}

func MakeConfigMapName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-proxy-configmap", c.GetName())
}
