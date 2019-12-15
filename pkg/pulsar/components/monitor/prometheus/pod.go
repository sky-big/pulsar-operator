package prometheus

import (
	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"

	"k8s.io/api/core/v1"
)

func makePodSpec(c *pulsarv1alpha1.PulsarCluster) v1.PodSpec {
	return v1.PodSpec{
		Containers:         []v1.Container{makeContainer(c)},
		Volumes:            makeVolumes(c),
		ServiceAccountName: MakeServiceAccountName(c),
	}
}

func makeContainer(c *pulsarv1alpha1.PulsarCluster) v1.Container {
	return v1.Container{
		Name:         "prometheus",
		Image:        pulsarv1alpha1.MonitorPrometheusImage,
		Ports:        makeContainerPort(c),
		VolumeMounts: makeContainerVolumeMount(c),
	}
}

func makeContainerPort(c *pulsarv1alpha1.PulsarCluster) []v1.ContainerPort {
	return []v1.ContainerPort{
		{
			Name:          "prometheus",
			ContainerPort: pulsarv1alpha1.PulsarPrometheusServerPort,
			Protocol:      v1.ProtocolTCP,
		},
	}
}

func makeContainerVolumeMount(c *pulsarv1alpha1.PulsarCluster) []v1.VolumeMount {
	return []v1.VolumeMount{
		{
			Name:      PrometheusDataVolumeName,
			MountPath: PrometheusDataVolumeMountPath,
		},
		{
			Name:      PrometheusConfigVolumeName,
			MountPath: PrometheusConfigVolumeMountPath,
		},
	}
}

func makeVolumes(c *pulsarv1alpha1.PulsarCluster) []v1.Volume {
	return []v1.Volume{
		{
			Name:         PrometheusDataVolumeName,
			VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}},
		},
		{
			Name:         PrometheusConfigVolumeName,
			VolumeSource: v1.VolumeSource{ConfigMap: &v1.ConfigMapVolumeSource{LocalObjectReference: v1.LocalObjectReference{Name: MakeConfigMapName(c)}}},
		},
	}
}
