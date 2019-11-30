package bookie

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"

	"k8s.io/api/core/v1"
)

func makePodSpec(c *pulsarv1alpha1.PulsarCluster) v1.PodSpec {
	p := v1.PodSpec{
		Affinity:       c.Spec.Zookeeper.Pod.Affinity,
		Containers:     []v1.Container{makeContainer(c)},
		InitContainers: []v1.Container{makeInitContainer(c)},
	}

	if isUseEmptyDirVolume(c) {
		p.Volumes = makeEmptyDirVolume(c)
	}

	return p
}

func makeContainer(c *pulsarv1alpha1.PulsarCluster) v1.Container {
	return v1.Container{
		Name:    "bookie",
		Image:   c.Spec.Bookie.Image.GenerateImage(),
		Command: makeContainerCommand(),
		Args:    makeContainerCommandArgs(),
		Ports:   makeContainerPort(c),
		Env:     makeContainerEnv(c),
		EnvFrom: makeContainerEnvFrom(c),

		VolumeMounts: []v1.VolumeMount{
			{
				Name:      makeJournalDataVolumeName(c),
				MountPath: BookieJournalDataMountPath,
			},
			{
				Name:      makeLedgersDataVolumeName(c),
				MountPath: BookieLedgersDataMountPath,
			},
		},

		ImagePullPolicy: c.Spec.Zookeeper.Image.PullPolicy,
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
			"bin/apply-config-from-env.py conf/pulsar_env.sh && " +
			"bin/pulsar bookie",
	}
}

func makeContainerPort(c *pulsarv1alpha1.PulsarCluster) []v1.ContainerPort {
	return []v1.ContainerPort{
		{
			Name:          "client",
			ContainerPort: pulsarv1alpha1.PulsarBookieServerPort,
			Protocol:      v1.ProtocolTCP,
		},
	}
}

func makeContainerEnv(c *pulsarv1alpha1.PulsarCluster) []v1.EnvVar {
	env := make([]v1.EnvVar, 0)
	return env
}

func makeContainerEnvFrom(c *pulsarv1alpha1.PulsarCluster) []v1.EnvFromSource {
	froms := make([]v1.EnvFromSource, 0)

	var configRef v1.ConfigMapEnvSource
	configRef.Name = MakeConfigMapName(c)

	froms = append(froms, v1.EnvFromSource{ConfigMapRef: &configRef})
	return froms
}

func makeInitContainer(c *pulsarv1alpha1.PulsarCluster) v1.Container {
	return v1.Container{
		Name:    "bookie-metaformat",
		Image:   fmt.Sprintf("%s:%s", pulsarv1alpha1.DefaultPulsarContainerRepository, pulsarv1alpha1.DefaultPulsarContainerVersion),
		Command: makeInitContainerCommand(),
		Args:    makeInitContainerCommandArgs(),
		EnvFrom: makeInitContainerEnvFrom(c),
	}
}

func makeInitContainerCommand() []string {
	return []string{
		"sh",
		"-c",
	}
}

func makeInitContainerCommandArgs() []string {
	return []string{
		"bin/apply-config-from-env.py conf/bookkeeper.conf && " +
			"bin/bookkeeper shell metaformat --nonInteractive || true;",
	}
}

func makeInitContainerEnvFrom(c *pulsarv1alpha1.PulsarCluster) []v1.EnvFromSource {
	froms := make([]v1.EnvFromSource, 0)

	var configRef v1.ConfigMapEnvSource
	configRef.Name = MakeConfigMapName(c)

	froms = append(froms, v1.EnvFromSource{ConfigMapRef: &configRef})
	return froms
}

func isUseEmptyDirVolume(c *pulsarv1alpha1.PulsarCluster) bool {
	return c.Spec.Bookie.StorageClassName == ""
}
