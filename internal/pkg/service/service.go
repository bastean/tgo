package service

import (
	"github.com/bastean/tgo/internal/pkg/service/cryptocurrency/coingecko"
	"github.com/bastean/tgo/internal/pkg/service/env"
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/logger/log"
	"github.com/bastean/tgo/internal/pkg/service/persistence/cassandra"
	"github.com/bastean/tgo/internal/pkg/service/user"
)

var Service = &struct {
	Cassandra, CoinGecko string
}{
	Cassandra: log.Service("Cassandra"),
	CoinGecko: log.Service("CoinGecko"),
}

var Module = &struct {
	User string
}{
	User: log.Module("User"),
}

var (
	err       error
	Cassandra *cassandra.Cassandra
	CoinGecko *coingecko.CoinGecko
)

func Up() error {
	log.EstablishingConnectionWith(Service.Cassandra)

	Cassandra, err = cassandra.Open(
		&cassandra.Auth{
			Hostname: env.DatabaseCassandraHostname,
			Port:     env.DatabaseCassandraPort,
			Username: env.DatabaseCassandraUser,
			Password: env.DatabaseCassandraPassword,
		},
		&cassandra.Config{
			Keyspace: env.DatabaseCassandraKeyspace,
			Strategy: "SimpleStrategy",
			Factor:   1,
		},
	)

	if err != nil {
		log.ConnectionFailedWith(Service.Cassandra)
		return errors.BubbleUp(err, "Up")
	}

	log.ConnectionEstablishedWith(Service.Cassandra)

	log.Starting(Service.CoinGecko)

	CoinGecko, err = coingecko.New(
		env.APICoinGeckoDemoKey,
		"https://api.coingecko.com/api/v3/",
	)

	if err != nil {
		log.CannotBeStarted(Service.CoinGecko)
		return errors.BubbleUp(err, "Up")
	}

	log.Started(Service.CoinGecko)

	log.Starting(Module.User)

	table, err := user.OpenTable(
		Cassandra,
		"users",
	)

	if err != nil {
		log.CannotBeStarted(Module.User)
		return errors.BubbleUp(err, "Up")
	}

	market := user.NewMarket(
		CoinGecko,
	)

	user.Start(
		table,
		market,
	)

	log.Started(Module.User)

	return nil
}

func Down() {
	log.ClosingConnectionWith(Service.Cassandra)

	cassandra.Close(Cassandra)

	log.ConnectionClosedWith(Service.Cassandra)
}
