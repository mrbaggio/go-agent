package newrelic

// DatastoreProduct encourages consistent metrics across New Relic agents.  You
// may create your own if your datastore is not listed below.
type DatastoreProduct string

// Datastore names used across New Relic agents:
const (
	DatastoreCassandra     DatastoreProduct = "Cassandra"
	DatastoreDerby         DatastoreProduct = "Derby"
	DatastoreElasticsearch DatastoreProduct = "Elasticsearch"
	DatastoreFirebird      DatastoreProduct = "Firebird"
	DatastoreIBMDB2        DatastoreProduct = "IBMDB2"
	DatastoreInformix      DatastoreProduct = "Informix"
	DatastoreMemcached     DatastoreProduct = "Memcached"
	DatastoreMongoDB       DatastoreProduct = "MongoDB"
	DatastoreMySQL         DatastoreProduct = "MySQL"
	DatastoreMSSQL         DatastoreProduct = "MSSQL"
	DatastoreOracle        DatastoreProduct = "Oracle"
	DatastorePostgres      DatastoreProduct = "Postgres"
	DatastoreRedis         DatastoreProduct = "Redis"
	DatastoreSolr          DatastoreProduct = "Solr"
	DatastoreSQLite        DatastoreProduct = "SQLite"
	DatastoreCouchDB       DatastoreProduct = "CouchDB"
	DatastoreRiak          DatastoreProduct = "Riak"
	DatastoreVoltDB        DatastoreProduct = "VoltDB"
)
