package v1alpha1

// Service
const (
	Service = "pulsar"
)

// Pulsar Cluster Phase Status List
const (
	// InitIng Phase
	PulsarClusterInitingPhase = "Initing"

	// Running Phase
	PulsarClusterRunningPhase = "Running"
)

const (
	// Zookeeper Component
	ZookeeperComponent = "zookeeper"

	// Broker Component
	BrokerComponent = "broker"

	// Bookie Component
	BookieComponent = "bookie"
)

const (
	// DefaultPulsarContainerRepository is the pulsar common container
	DefaultPulsarContainerRepository = "apachepulsar/pulsar"

	// DefaultPulsarContainerVersion is the default tag used for container
	DefaultPulsarContainerVersion = "latest"

	// DefaultAllPulsarContainerRepository is the default docker repo for components
	DefaultAllPulsarContainerRepository = "apachepulsar/pulsar-all"

	// DefaultAllPulsarContainerVersion is the default tag used for components
	DefaultAllPulsarContainerVersion = "latest"

	// DefaultZkContainerPolicy is the default container pull policy used
	DefaultContainerPolicy = "Always"
)

// Labels
const (
	// App
	LabelService = "app"

	// Cluster
	LabelCluster = "cluster"

	// Component
	LabelComponent = "component"
)

func MakeLabels(c *PulsarCluster, component string) map[string]string {
	labels := make(map[string]string)
	labels[LabelService] = Service
	labels[LabelCluster] = c.GetName()
	labels[LabelComponent] = component
	return labels
}
