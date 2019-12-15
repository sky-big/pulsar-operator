package dashboard

import (
	"fmt"
	broker2 "github.com/sky-big/pulsar-operator/pkg/pulsar/components/broker"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"k8s.io/api/core/v1"
)

func makePodSpec(c *pulsarv1alpha1.PulsarCluster) v1.PodSpec {
	return v1.PodSpec{
		Containers: []v1.Container{makeContainer(c)},
	}
}

func makeContainer(c *pulsarv1alpha1.PulsarCluster) v1.Container {
	return v1.Container{
		Name:  "dashboard",
		Image: pulsarv1alpha1.MonitorDashboardImage,
		Ports: makeContainerPort(c),
		Env:   makeContainerEnv(c),
	}
}

func makeContainerPort(c *pulsarv1alpha1.PulsarCluster) []v1.ContainerPort {
	return []v1.ContainerPort{
		{
			Name:          "dashboard",
			ContainerPort: pulsarv1alpha1.PulsarDashboardServerPort,
			Protocol:      v1.ProtocolTCP,
		},
	}
}

func makeContainerEnv(c *pulsarv1alpha1.PulsarCluster) []v1.EnvVar {
	brokerUrl := fmt.Sprintf("http://%s:%d/", broker2.MakeServiceName(c), pulsarv1alpha1.PulsarBrokerHttpServerPort)
	env := make([]v1.EnvVar, 0)
	broker := v1.EnvVar{
		Name:  "SERVICE_URL",
		Value: brokerUrl,
	}
	env = append(env, broker)
	return env
}
