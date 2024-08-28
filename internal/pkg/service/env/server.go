package env

import (
	"os"
)

var (
	ServerEchoHostname string
	ServerEchoPort     string
	ServerEchoURL      string
)

func Server() {
	ServerEchoHostname = os.Getenv("TGO_SERVER_ECHO_HOSTNAME")
	ServerEchoPort = os.Getenv("TGO_SERVER_ECHO_PORT")
	ServerEchoURL = os.Getenv("TGO_SERVER_ECHO_URL")
}

func HasServerEchoProxy() (string, bool) {
	proxy := os.Getenv("TGO_DEV_AIR_PROXY_PORT")

	if proxy != "" && proxy != ServerEchoPort {
		return proxy, true
	}

	return "", false
}
