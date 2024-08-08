package cassandra

import (
	"fmt"

	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/gocql/gocql"
)

type Auth struct {
	Hostname, Port, Username, Password string
}

type Config struct {
	Keyspace, Strategy string
	Factor             int
}

type Cassandra struct {
	*gocql.Session
	Keyspace string
}

func Open(auth *Auth, config *Config) (*Cassandra, error) {
	cluster := gocql.NewCluster(fmt.Sprintf("%s:%s", auth.Hostname, auth.Port))

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: auth.Username,
		Password: auth.Password,
	}

	session, err := cluster.CreateSession()

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Open",
			What:  "Failure to create a Cassandra session",
			Who:   err,
		})
	}

	create := fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': '%s', 'replication_factor': %d}", config.Keyspace, config.Strategy, config.Factor)

	err = session.Query(create).Exec()

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Open",
			What:  "Failure to create a Keyspace",
			Why: errors.Meta{
				"Keyspace": config.Keyspace,
			},
			Who: err,
		})
	}

	return &Cassandra{
		Session:  session,
		Keyspace: config.Keyspace,
	}, nil
}

func Close(session *Cassandra) {
	session.Close()
}
