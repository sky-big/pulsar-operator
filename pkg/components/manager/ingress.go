package manager

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"

	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func MakeIngress(c *pulsarv1alpha1.PulsarCluster) *v1beta1.Ingress {
	return &v1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        MakeIngressName(c),
			Namespace:   c.Namespace,
			Labels:      pulsarv1alpha1.MakeComponentLabels(c, pulsarv1alpha1.ManagerComponent),
			Annotations: c.Spec.Manager.Annotations,
		},
		Spec: makeIngressSpec(c),
	}
}

func MakeIngressName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-manager-ingress", c.GetName())
}

func makeIngressSpec(c *pulsarv1alpha1.PulsarCluster) v1beta1.IngressSpec {
	s := v1beta1.IngressSpec{
		Rules: make([]v1beta1.IngressRule, 0),
	}

	if c.Spec.Manager.Host != "" {
		s.Rules = append(s.Rules, makeRule(c))
	}
	return s
}

func makeRule(c *pulsarv1alpha1.PulsarCluster) v1beta1.IngressRule {
	r := v1beta1.IngressRule{
		Host: c.Spec.Manager.Host,
		IngressRuleValue: v1beta1.IngressRuleValue{
			HTTP: &v1beta1.HTTPIngressRuleValue{
				Paths: make([]v1beta1.HTTPIngressPath, 0),
			},
		},
	}
	path := v1beta1.HTTPIngressPath{
		Path: "/",
		Backend: v1beta1.IngressBackend{
			ServiceName: MakeServiceName(c),
			ServicePort: intstr.FromInt(pulsarv1alpha1.PulsarManagerServerPort),
		},
	}
	r.IngressRuleValue.HTTP.Paths = append(r.IngressRuleValue.HTTP.Paths, path)
	return r
}
