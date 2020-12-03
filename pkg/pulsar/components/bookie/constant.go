package bookie

const (
	BookieMemData = "\" -Xms64m -Xmx256m -XX:MaxDirectMemorySize=256m\""

	StatsProviderClass = "org.apache.bookkeeper.stats.prometheus.PrometheusMetricsProvider"

	BookieJournalDataMountPath = "/pulsar/data/bookkeeper/journal"

	BookieLedgersDataMountPath = "/pulsar/data/bookkeeper/ledgers"
)

// Annotations
var StatefulSetAnnotations map[string]string

func init() {
	StatefulSetAnnotations = make(map[string]string)
	StatefulSetAnnotations["prometheus.io/scrape"] = "true"
	StatefulSetAnnotations["prometheus.io/port"] = "8000"
}
