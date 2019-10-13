package v1alpha1

const (
	// Pulsar cluster active flag
	MonitorActived = "true"

	// Dashboard image
	MonitorDashboardImage = "apachepulsar/pulsar-dashboard:latest"

	// Prometheus image
	MonitorPrometheusImage = "prom/prometheus:v1.6.3"

	// Grafana image
	MonitorGrafanaImage = "apachepulsar/pulsar-grafana:latest"
)

// Monitor defines the desired state of Monitor
// +k8s:openapi-gen=true
type Monitor struct {
	// Is active pulsar cluster monitor flag.
	IsActive string `json:"isActive,omitempty"`

	// DashboardPort (DEPRECATED) is the expected port of the pulsar dashboard.
	DashboardPort int32 `json:"dashboardPort,omitempty"`

	// PrometheusPort (DEPRECATED) is the expected port of the pulsar prometheus.
	PrometheusPort int32 `json:"prometheusPort,omitempty"`

	// GrafanaPort (DEPRECATED) is the expected port of the pulsar grafana.
	GrafanaPort int32 `json:"grafanaPort,omitempty"`
}

func (m *Monitor) SetDefault(cluster *PulsarCluster) bool {
	changed := false

	if m.IsActive == "" {
		m.IsActive = "false"
		changed = true
	}
	return changed
}
