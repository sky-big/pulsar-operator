package prometheus

const (
	PulsarPrometheusContainerPort = 9090

	PrometheusDataVolumeName = "data-volume"

	PrometheusDataVolumeMountPath = "/prometheus"

	PrometheusConfigVolumeName = "config-volume"

	PrometheusConfigVolumeMountPath = "/etc/prometheus"
)
