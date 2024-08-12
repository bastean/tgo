package cassandra

import (
	"github.com/bastean/tgo/pkg/context/infrastructure/persistence/cassandra"
)

type (
	Cassandra = cassandra.Cassandra
	Auth      = cassandra.Auth
	Config    = cassandra.Config
)

var (
	Open  = cassandra.Open
	Close = cassandra.Close
)
