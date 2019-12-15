package grafana

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/monitor/prometheus"

	"k8s.io/api/core/v1"
)

func makePodSpec(c *pulsarv1alpha1.PulsarCluster) v1.PodSpec {
	return v1.PodSpec{
		Containers: []v1.Container{makeContainer(c)},
	}
}

func makeContainer(c *pulsarv1alpha1.PulsarCluster) v1.Container {
	return v1.Container{
		Name:  "grafana",
		Image: pulsarv1alpha1.MonitorGrafanaImage,
		Ports: makeContainerPort(c),
		Env:   makeContainerEnv(c),
	}
}

func makeContainerPort(c *pulsarv1alpha1.PulsarCluster) []v1.ContainerPort {
	return []v1.ContainerPort{
		{
			Name:          "grafana",
			ContainerPort: pulsarv1alpha1.PulsarGrafanaServerPort,
			Protocol:      v1.ProtocolTCP,
		},
	}
}

func makeContainerEnv(c *pulsarv1alpha1.PulsarCluster) []v1.EnvVar {
	prometheusUrl := fmt.Sprintf("http://%s:%d/", prometheus.MakeServiceName(c), pulsarv1alpha1.PulsarPrometheusServerPort)
	env := make([]v1.EnvVar, 0)
	p := v1.EnvVar{
		Name:  "PROMETHEUS_URL",
		Value: prometheusUrl,
	}
	env = append(env, p)
	return env
}
