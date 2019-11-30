package manager

import (
	"strconv"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"

	"k8s.io/api/core/v1"
)

func makePodSpec(c *pulsarv1alpha1.PulsarCluster) v1.PodSpec {
	return v1.PodSpec{
		Containers: []v1.Container{makeContainer(c)},
		Volumes:    makeVolumes(c),
	}
}

func makeContainer(c *pulsarv1alpha1.PulsarCluster) v1.Container {
	return v1.Container{
		Name:         "manager",
		Image:        c.Spec.Manager.Image.GenerateImage(),
		Ports:        makeContainerPort(c),
		VolumeMounts: makeContainerVolumeMount(c),
		Env:          makeContainerEnv(c),
	}
}

func makeContainerPort(c *pulsarv1alpha1.PulsarCluster) []v1.ContainerPort {
	return []v1.ContainerPort{
		{
			Name:          "manager",
			ContainerPort: pulsarv1alpha1.PulsarManagerServerPort,
			Protocol:      v1.ProtocolTCP,
		},
	}
}

func makeContainerVolumeMount(c *pulsarv1alpha1.PulsarCluster) []v1.VolumeMount {
	return []v1.VolumeMount{
		{
			Name:      PulsarManagerVolumeName,
			MountPath: PulsarManagerVolumeMountPath,
		},
	}
}

func makeVolumes(c *pulsarv1alpha1.PulsarCluster) []v1.Volume {
	return []v1.Volume{
		{
			Name:         PulsarManagerVolumeName,
			VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}},
		},
	}
}

func makeContainerEnv(c *pulsarv1alpha1.PulsarCluster) []v1.EnvVar {
	env := make([]v1.EnvVar, 0)

	// pulsar cluster
	cluster := v1.EnvVar{
		Name:  "PULSAR_CLUSTER",
		Value: c.GetName(),
	}
	env = append(env, cluster)

	// redirect host
	redirect := v1.EnvVar{
		Name:  "REDIRECT_HOST",
		Value: "http://127.0.0.1",
	}
	env = append(env, redirect)

	// redirect port
	redirectPort := v1.EnvVar{
		Name:  "REDIRECT_PORT",
		Value: strconv.Itoa(pulsarv1alpha1.PulsarManagerServerPort),
	}
	env = append(env, redirectPort)

	// driver
	driver := v1.EnvVar{
		Name:  "DRIVER_CLASS_NAME",
		Value: "org.postgresql.Driver",
	}
	env = append(env, driver)

	// url
	url := v1.EnvVar{
		Name:  "URL",
		Value: "jdbc:postgresql://127.0.0.1:5432/pulsar_manager",
	}
	env = append(env, url)

	// user name
	userName := v1.EnvVar{
		Name:  "USERNAME",
		Value: c.Spec.Manager.UserName,
	}
	env = append(env, userName)

	// user password
	userPassword := v1.EnvVar{
		Name:  "PASSWORD",
		Value: c.Spec.Manager.UserPassword,
	}
	env = append(env, userPassword)

	// log level
	logLevel := v1.EnvVar{
		Name:  "LOG_LEVEL",
		Value: "DEBUG",
	}
	env = append(env, logLevel)

	return env
}
