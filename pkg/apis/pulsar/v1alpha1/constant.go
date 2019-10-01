package v1alpha1

// Service
const (
	Service = "pulsar"
)

// Pulsar Cluster Phase Status List
const (
	// InitIng Phase
	PulsarClusterInitingPhase = "Initing"

	// Launching Phase
	PulsarClusterLaunchingPhase = "Launching"

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

	// Bookie AutoRecovery Child Component
	BookieAutoRecoveryComponent = "bookie-autorecovery"
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

	// ChildComponent
	LabelChildComponent = "child-component"
)

func MakeLabels(c *PulsarCluster, component string) map[string]string {
	return MakeAllLabels(c, component, "")
}

func MakeAllLabels(c *PulsarCluster, component string, childComponent string) map[string]string {
	labels := make(map[string]string)
	labels[LabelService] = Service
	labels[LabelCluster] = c.GetName()
	labels[LabelComponent] = component
	if childComponent != "" {
		labels[LabelChildComponent] = childComponent
	}
	return labels
}

// Service
const (
	// Service Domain
	ServiceDomain = "svc.cluster.local"
)

// Default Number
const (
	// zookeeper number default num is 3
	ZookeeperClusterDefaultNodeNum = 3

	// broker number default num is 3
	BrokerClusterDefaultNodeNum = 3

	// bookie number default num is 3
	BookieClusterDefaultNodeNum = 3

	// bookie autorecovery default num is 3
	BookieAutoRecoveryClusterDefaultNodeNum = 3
)
