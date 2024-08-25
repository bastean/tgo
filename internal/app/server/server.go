package server

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/bastean/tgo/internal/app/server/router"
	"github.com/bastean/tgo/internal/pkg/service/env"
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/logger/log"
)

var Server = &struct {
	Echo string
}{
	Echo: log.Server("Echo"),
}

var (
	//go:embed static
	Files embed.FS
	App   *http.Server
)

func Up() error {
	log.Starting(Server.Echo)

	App = &http.Server{
		Addr:    ":" + env.ServerEchoPort,
		Handler: router.New(&Files),
	}

	log.Started(Server.Echo)

	log.Info(fmt.Sprintf("%s listening on %s", Server.Echo, env.ServerEchoURL))

	if proxy, ok := env.HasServerEchoProxy(); ok {
		log.Info(fmt.Sprintf("%s proxy listening on %s", Server.Echo, strings.Replace(env.ServerEchoURL, env.ServerEchoPort, proxy, 1)))
	}

	if err := App.ListenAndServe(); errors.IsNot(err, http.ErrServerClosed) {
		log.CannotBeStarted(Server.Echo)

		return errors.NewInternal(&errors.Bubble{
			Where: "Up",
			What:  "Failure to start Server",
			Who:   err,
		})
	}

	return nil
}

func Down(ctx context.Context) error {
	log.Stopping(Server.Echo)

	if err := App.Shutdown(ctx); err != nil {
		log.CannotBeStopped(Server.Echo)

		return errors.NewInternal(&errors.Bubble{
			Where: "Down",
			What:  "Failure to shutdown Server",
			Who:   err,
		})
	}

	log.Stopped(Server.Echo)

	return nil
}
