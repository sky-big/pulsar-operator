package proxy

import (
	"fmt"
	broker2 "github.com/sky-big/pulsar-operator/pkg/pulsar/components/broker"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeConfigMap(c *pulsarv1alpha1.PulsarCluster) *v1.ConfigMap {
	brokerServiceName := broker2.MakeServiceName(c)
	webServiceUrl := fmt.Sprintf("http://%s.%s.%s:%d",
		brokerServiceName, c.Namespace, pulsarv1alpha1.ServiceDomain, pulsarv1alpha1.PulsarBrokerHttpServerPort)
	brokerServiceUrl := fmt.Sprintf("pulsar://%s.%s.%s:%d",
		brokerServiceName, c.Namespace, pulsarv1alpha1.ServiceDomain, pulsarv1alpha1.PulsarBrokerPulsarServerPort)

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
