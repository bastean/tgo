package env

import (
	"os"
)

var (
	APICoinGeckoDemoKey string
)

func API() {
	APICoinGeckoDemoKey = os.Getenv("TGO_API_COINGECKO_DEMO_KEY")
}
