package ingress

import (
	"fmt"

	pulsarv1alpha1 "github.com/sky-big/pulsar-operator/pkg/apis/pulsar/v1alpha1"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/monitor/grafana"
	"github.com/sky-big/pulsar-operator/pkg/pulsar/components/monitor/prometheus"

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
			Labels:      pulsarv1alpha1.MakeComponentLabels(c, pulsarv1alpha1.MonitorComponent),
			Annotations: c.Spec.Monitor.Ingress.Annotations,
		},
		Spec: makeIngressSpec(c),
	}
}

func MakeIngressName(c *pulsarv1alpha1.PulsarCluster) string {
	return fmt.Sprintf("%s-monitor-ingress", c.GetName())
}

func makeIngressSpec(c *pulsarv1alpha1.PulsarCluster) v1beta1.IngressSpec {
	s := v1beta1.IngressSpec{
		Rules: make([]v1beta1.IngressRule, 0),
	}

	if c.Spec.Monitor.Grafana.Host != "" {
		s.Rules = append(s.Rules, makeGrafanaRule(c))
	}

	if c.Spec.Monitor.Prometheus.Host != "" {
		s.Rules = append(s.Rules, makePrometheusRule(c))
	}
	return s
}

func makeGrafanaRule(c *pulsarv1alpha1.PulsarCluster) v1beta1.IngressRule {
	r := v1beta1.IngressRule{
		Host: c.Spec.Monitor.Grafana.Host,
		IngressRuleValue: v1beta1.IngressRuleValue{
			HTTP: &v1beta1.HTTPIngressRuleValue{
				Paths: make([]v1beta1.HTTPIngressPath, 0),
			},
		},
	}
	path := v1beta1.HTTPIngressPath{
		Path: "/",
		Backend: v1beta1.IngressBackend{
			ServiceName: grafana.MakeServiceName(c),
			ServicePort: intstr.FromInt(pulsarv1alpha1.PulsarGrafanaServerPort),
		},
	}
	r.IngressRuleValue.HTTP.Paths = append(r.IngressRuleValue.HTTP.Paths, path)
	return r
}

func makePrometheusRule(c *pulsarv1alpha1.PulsarCluster) v1beta1.IngressRule {
	r := v1beta1.IngressRule{
		Host: c.Spec.Monitor.Prometheus.Host,
		IngressRuleValue: v1beta1.IngressRuleValue{
			HTTP: &v1beta1.HTTPIngressRuleValue{
				Paths: make([]v1beta1.HTTPIngressPath, 0),
			},
		},
	}
	path := v1beta1.HTTPIngressPath{
		Path: "/",
		Backend: v1beta1.IngressBackend{
			ServiceName: prometheus.MakeServiceName(c),
			ServicePort: intstr.FromInt(pulsarv1alpha1.PulsarPrometheusServerPort),
		},
	}
	r.IngressRuleValue.HTTP.Paths = append(r.IngressRuleValue.HTTP.Paths, path)
	return r
}
