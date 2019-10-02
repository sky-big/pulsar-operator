package v1alpha1

// service
const (
	Service = "pulsar"
)

// pulsar cluster phase
const (
	// Initing phase
	PulsarClusterInitingPhase = "Initing"

	// Launching phase
	PulsarClusterLaunchingPhase = "Launching"

	// Running phase
	PulsarClusterRunningPhase = "Running"
)

// pulsar cluster components
const (
	// Zookeeper component
	ZookeeperComponent = "zookeeper"

	// Broker component
	BrokerComponent = "broker"

	// Bookie component
	BookieComponent = "bookie"

	// Bookie AutoRecovery child component
	BookieAutoRecoveryComponent = "bookie-autorecovery"

	// Proxy component
	ProxyComponent = "proxy"
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

// service
const (
	// service domain
	ServiceDomain = "svc.cluster.local"
)

// default number
const (
	// Zookeeper number default num is 3
	ZookeeperClusterDefaultNodeNum = 3

	// Broker number default num is 3
	BrokerClusterDefaultNodeNum = 3

	// Bookie number default num is 3
	BookieClusterDefaultNodeNum = 3

	// Bookie Autorecovery default num is 3
	BookieAutoRecoveryClusterDefaultNodeNum = 3

	// Proxy number default num is 3
	ProxyClusterDefaultNodeNum = 3
)

// All component ports
const (
	// Container client default port
	ZookeeperContainerClientDefaultPort = 2181

	// Container server default port
	ZookeeperContainerServerDefaultPort = 2888

	// Container leader election port
	ZookeeperContainerLeaderElectionPort = 3888

	// Broker service port
	PulsarBrokerPulsarServicePort = 6650

	// Broker http service port
	PulsarBrokerHttpServicePort = 8080

	// Bookie service port
	BookieServerPort = 3181
)
