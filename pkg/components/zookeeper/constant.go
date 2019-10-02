package zookeeper

const (
	// Mem
	PulsarMemData = "\" -Xms100m -Xmx256m \""

	// GC
	PulsarGCData = "\" -XX:+UseG1GC -XX:MaxGCPauseMillis=10\""

	// Container zookeeper server list env key
	ContainerZookeeperServerList = "ZOOKEEPER_SERVERS"

	// Container data volume name
	ContainerDataVolumeName = "datadir"

	// Container zookeeper data path
	ContainerDataPath = "/pulsar/data"

	// ReadinessProbe script
	ReadinessProbeScript = "bin/pulsar-zookeeper-ruok.sh"

	// LivenessProbe script
	LivenessProbeScript = "bin/pulsar-zookeeper-ruok.sh"
)
