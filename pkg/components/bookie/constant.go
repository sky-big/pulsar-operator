package bookie

const (
	PulsarMemData = "\" -Xms64m -Xmx256m -XX:MaxDirectMemorySize=256m\""

	DbStorage_writeCacheMaxSizeMb = "32"

	DbStorage_readAheadCacheMaxSizeMb = "32"

	StatsProviderClass = "org.apache.bookkeeper.stats.prometheus.PrometheusMetricsProvider"

	BookieServerPort = 3181

	BookieJournalDataMountPath = "/pulsar/data/bookkeeper/journal"

	BookieLedgersDataMountPath = "/pulsar/data/bookkeeper/ledgers"
)
