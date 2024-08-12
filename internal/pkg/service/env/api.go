package env

import (
	"os"
)

var (
	APICoinGeckoDemoKey = os.Getenv("TGO_API_COINGECKO_DEMO_KEY")
)
