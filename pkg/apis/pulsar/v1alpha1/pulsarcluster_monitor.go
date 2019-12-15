package v1alpha1

const (
	// Prometheus image
	MonitorPrometheusImage = "prom/prometheus:v1.6.3"

	// Grafana image
	MonitorGrafanaImage = "apachepulsar/pulsar-grafana:latest"
)

// Monitor defines the desired state of Monitor
// +k8s:openapi-gen=true
type Monitor struct {
	// Is enable pulsar cluster monitor flag.
	Enable bool `json:"enable,omitempty"`

	// Prometheus
	Prometheus Prometheus `json:"prometheus,omitempty"`

	// Grafana
	Grafana Grafana `json:"grafana,omitempty"`

	// Ingress
	Ingress MonitorIngress `json:"ingress,omitempty"`
}

// Pulsar cluster prometheus spec
// +k8s:openapi-gen=true
type Prometheus struct {
	// Host (DEPRECATED) is the expected host of the pulsar prometheus.
	Host string `json:"host,omitempty"`

	// NodePort (DEPRECATED) is the expected port of the pulsar prometheus.
	NodePort int32 `json:"nodePort,omitempty"`
}

// Pulsar cluster grafana spec
// +k8s:openapi-gen=true
type Grafana struct {
	// Host (DEPRECATED) is the expected host of the pulsar grafana.
	Host string `json:"host,omitempty"`

	// NodePort (DEPRECATED) is the expected port of the pulsar grafana.
	NodePort int32 `json:"nodePort,omitempty"`
}

// MonitorIngress defines the pulsar cluster exposed
// +k8s:openapi-gen=true
type MonitorIngress struct {
	// enable ingress
	Enable bool `json:"enable,omitempty"`

	// Ingress additional annotation
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (m *Monitor) SetDefault(cluster *PulsarCluster) bool {
	return false
}
