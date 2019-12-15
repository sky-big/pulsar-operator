package proxy

const (
	PulsarMemData = "\" -Xms4g -Xmx4g -XX:MaxDirectMemorySize=4g\""
)

// Annotations
var DeploymentAnnotations map[string]string

func init() {
	DeploymentAnnotations = make(map[string]string)
	DeploymentAnnotations["prometheus.io/scrape"] = "true"
	DeploymentAnnotations["prometheus.io/port"] = "8080"
}
