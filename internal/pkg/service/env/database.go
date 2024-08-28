package env

import (
	"os"
)

var (
	DatabaseCassandraHostname string
	DatabaseCassandraPort     string
	DatabaseCassandraUser     string
	DatabaseCassandraPassword string
	DatabaseCassandraKeyspace string
)

func Database() {
	DatabaseCassandraHostname = os.Getenv("DATABASE_CASSANDRA_HOSTNAME")
	DatabaseCassandraPort = os.Getenv("DATABASE_CASSANDRA_PORT")
	DatabaseCassandraUser = os.Getenv("DATABASE_CASSANDRA_USER")
	DatabaseCassandraPassword = os.Getenv("DATABASE_CASSANDRA_PASSWORD")
	DatabaseCassandraKeyspace = os.Getenv("DATABASE_CASSANDRA_KEYSPACE")
}
