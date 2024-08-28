package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bastean/tgo/internal/app/bot"
	"github.com/bastean/tgo/internal/app/cli"
	"github.com/bastean/tgo/internal/app/server"
	"github.com/bastean/tgo/internal/pkg/service"
	"github.com/bastean/tgo/internal/pkg/service/logger/log"
)

var (
	err error
)

var (
	Services = "Services"
	Apps     = "Apps"
)

func main() {
	if err = cli.Up(); err != nil {
		log.Fatal(err.Error())
	}

	log.Logo()

	log.Starting(Services)

	if err = service.Up(); err != nil {
		log.Fatal(err.Error())
	}

	log.Started(Services)

	log.Starting(Apps)

	go func() {
		if err := server.Up(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	go func() {
		if err := bot.Up(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Started(Apps)

	log.Info("Press Ctrl+C to exit")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	log.Stopping(Apps)

	errServer := server.Down(ctx)

	errBot := bot.Down()

	if err := errors.Join(errServer, errBot); err != nil {
		log.Error(err.Error())
	}

	log.Stopped(Apps)

	log.Stopping(Services)

	service.Down()

	log.Stopped(Services)

	<-ctx.Done()

	log.Info("Exiting...")
}
