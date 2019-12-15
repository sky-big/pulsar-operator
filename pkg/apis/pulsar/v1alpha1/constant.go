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

	// Proxy component
	ProxyComponent = "proxy"

	// Monitor component
	MonitorComponent = "monitor"

	// Manager component
	ManagerComponent = "manager"
)

// bookie child component
const (
	// Bookie AutoRecovery child component
	BookieAutoRecoveryComponent = "bookie-autorecovery"
)

// monitor child component
const (
	// Monitor prometheus component
	MonitorPrometheusComponent = "monitor-prometheus"

	// Monitor grafana component
	MonitorGrafanaComponent = "monitor-grafana"
)

const (
	// DefaultPulsarManagerRepository is default docker image name of pulsar pulsar
	DefaultPulsarContainerRepository = "apachepulsar/pulsar"

	// DefaultPulsarContainerVersion is the default tag used for container
	DefaultPulsarContainerVersion = "latest"

	// DefaultAllPulsarContainerRepository is the default docker repo for components
	DefaultAllPulsarContainerRepository = "apachepulsar/pulsar-all"

	// DefaultAllPulsarContainerVersion is the default tag used for components
	DefaultAllPulsarContainerVersion = "latest"

	// DefaultPulsarManagerRepository is default docker image name of pulsar manager
	DefaultPulsarManagerContainerRepository = "apachepulsar/pulsar-manager"

	// DefaultPulsarManagerContainerVersion is
	DefaultPulsarManagerContainerVersion = "v0.1.0"

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

func MakeComponentLabels(c *PulsarCluster, component string) map[string]string {
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

	// Broker server port
	PulsarBrokerPulsarServerPort = 6650

	// Broker http server port
	PulsarBrokerHttpServerPort = 8080

	// Bookie server port
	PulsarBookieServerPort = 3181

	// Grafana server port
	PulsarGrafanaServerPort = 3000

	// Prometheus server port
	PulsarPrometheusServerPort = 9090

	// Manager server port
	PulsarManagerServerPort = 9527
)

// Storage default capacity
const (
	// journal storage default capacity
	JournalStorageDefaultCapacity = 1

	// ledgers storage default capacity
	LedgersStorageDefaultCapacity = 10
)

// Manager component default user name and user password
const (
	// manager user default name
	ManagerDefaultUserName = "pulsar"

	// manager user default password
	ManagerDefaultUserPassword = "pulsar"
)
