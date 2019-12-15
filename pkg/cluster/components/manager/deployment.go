package manager

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeDeployment(c *pulsarv1alpha1.PulsarCluster) *appsv1.Deployment {
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      MakeDeploymentName(c),
			Namespace: c.Namespace,
			Labels:    pulsarv1alpha1.MakeComponentLabels(c, pulsarv1alpha1.ManagerComponent),
		},
		Spec: makeDeploymentSpec(c),
	}
}

func MakeDeploymentName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-manager-deployment", c.GetName())
}

func makeDeploymentSpec(c *pulsarv1alpha1.PulsarCluster) appsv1.DeploymentSpec {
	var replica int32 = 1
	return appsv1.DeploymentSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: pulsarv1alpha1.MakeComponentLabels(c, pulsarv1alpha1.ManagerComponent),
		},
		Replicas: &replica,
		Template: makeDeploymentPodTemplate(c),
	}
}

func makeDeploymentPodTemplate(c *pulsarv1alpha1.PulsarCluster) v1.PodTemplateSpec {
	return v1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: c.GetName(),
			Labels:       pulsarv1alpha1.MakeComponentLabels(c, pulsarv1alpha1.ManagerComponent),
		},
		Spec: makePodSpec(c),
	}
}
