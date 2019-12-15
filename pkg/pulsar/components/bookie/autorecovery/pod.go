package autorecovery

import (
	"fmt"
	bookie2 "github.com/sky-big/pulsar-operator/pkg/pulsar/components/bookie"

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
		Name:    "replication-worker",
		Image:   fmt.Sprintf("%s:%s", pulsarv1alpha1.DefaultPulsarContainerRepository, pulsarv1alpha1.DefaultPulsarContainerVersion),
		Command: makeContainerCommand(),
		Args:    makeContainerCommandArgs(),
		Env:     makeContainerEnv(c),
		EnvFrom: makeContainerEnvFrom(c),
	}
}

func makeContainerCommand() []string {
	return []string{
		"sh",
		"-c",
	}
}

func makeContainerCommandArgs() []string {
	return []string{
		"bin/apply-config-from-env.py conf/bookkeeper.conf && " +
			"bin/bookkeeper autorecovery",
	}
}

func makeContainerEnv(c *pulsarv1alpha1.PulsarCluster) []v1.EnvVar {
	env := make([]v1.EnvVar, 0)
	env = append(env,
		v1.EnvVar{
			Name:  "PULSAR_MEM",
			Value: PulsarMemData,
		},
		v1.EnvVar{
			Name:  "PULSAR_GC",
			Value: PulsarGCData,
		},
	)
	return env
}

func makeContainerEnvFrom(c *pulsarv1alpha1.PulsarCluster) []v1.EnvFromSource {
	froms := make([]v1.EnvFromSource, 0)

	var configRef v1.ConfigMapEnvSource
	configRef.Name = bookie2.MakeConfigMapName(c)

	froms = append(froms, v1.EnvFromSource{ConfigMapRef: &configRef})
	return froms
}
