package broker

const (
	PulsarMemData = "\" -Xms64m -Xmx128m -XX:MaxDirectMemorySize=128m\""

	ManagedLedgerDefaultEnsembleSize = "1"

	ManagedLedgerDefaultWriteQuorum = "1"

	ManagedLedgerDefaultAckQuorum = "1"

	FunctionsWorkerEnabled = "true"

	AdvertisedAddress = "advertisedAddress"
)

// Annotations
var DeploymentAnnotations map[string]string

func init() {
	DeploymentAnnotations = make(map[string]string)
	DeploymentAnnotations["prometheus.io/scrape"] = "true"
	DeploymentAnnotations["prometheus.io/port"] = "8080"
}
