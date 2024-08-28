package cli

import (
	"flag"
	"fmt"

	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/logger/log"
	"github.com/joho/godotenv"
)

const (
	cli = "tgo"
)

var (
	env string
)

func usage() {
	log.Logo()

	fmt.Print("Example of interoperability between a Web App, a Telegram Bot and a third-party API.\n\n")

	fmt.Printf("Usage: %s [flags]\n\n", cli)

	flag.PrintDefaults()
}

func Up() error {
	flag.StringVar(&env, "env", "", "Path to ENV file (required)")

	flag.Usage = usage

	flag.Parse()

	if err := godotenv.Load(env); err != nil && env != "" {
		return errors.NewInternal(&errors.Bubble{
			Where: "Up",
			What:  "Failure to load ENV file",
			Who:   err,
		})
	}

	return nil
}
