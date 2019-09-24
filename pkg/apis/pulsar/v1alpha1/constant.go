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
	// DefaultZkContainerRepository is the default docker repo for the container
	DefaultContainerRepository = "apachepulsar/pulsar-all"

	// DefaultZkContainerVersion is the default tag used for for the container
	DefaultContainerVersion = "latest"

	// DefaultZkContainerPolicy is the default container pull policy used
	DefaultContainerPolicy = "Always"
)

// Zookeeper
const (
	// Container Client Default Port
	ZookeeperContainerClientDefaultPort = 2181

	// Container Server Default Port
	ZookeeperContainerServerDefaultPort = 2888

	// Container Leader Election Port
	ZookeeperContainerLeaderElectionPort = 3888
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
