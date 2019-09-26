package zookeeper

const (
	// Mem
	PulsarMemData = "\" -Xms100m -Xmx256m \""

	// GC
	PulsarGCData = "\" -XX:+UseG1GC -XX:MaxGCPauseMillis=10\""

	// Container Zookeeper Server List Env Key
	ContainerZookeeperServerList = "ZOOKEEPER_SERVERS"

	// Container Data Volume Name
	ContainerDataVolumeName = "datadir"

	// Container Zookeeper Data Path
	ContainerDataPath = "/pulsar/data"

	// ReadinessProbe Script
	ReadinessProbeScript = "bin/pulsar-zookeeper-ruok.sh"

	// LivenessProbe Script
	LivenessProbeScript = "bin/pulsar-zookeeper-ruok.sh"

	// Container Client Default Port
	ZookeeperContainerClientDefaultPort = 2181

	// Container Server Default Port
	ZookeeperContainerServerDefaultPort = 2888

	// Container Leader Election Port
	ZookeeperContainerLeaderElectionPort = 3888
)
