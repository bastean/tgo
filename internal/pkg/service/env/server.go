package env

import (
	"os"
)

var (
	ServerEchoHostname          = os.Getenv("TGO_SERVER_ECHO_HOSTNAME")
	ServerEchoPort              = os.Getenv("TGO_SERVER_ECHO_PORT")
	ServerEchoURL               = os.Getenv("TGO_SERVER_ECHO_URL")
	ServerEchoMode              = os.Getenv("TGO_SERVER_ECHO_MODE")
	ServerEchoAllowedHosts      = os.Getenv("TGO_SERVER_ECHO_ALLOWED_HOSTS")
	ServerEchoCookieSecretKey   = os.Getenv("TGO_SERVER_ECHO_COOKIE_SECRET_KEY")
	ServerEchoCookieSessionName = os.Getenv("TGO_SERVER_ECHO_COOKIE_SESSION_NAME")
)

func HasServerEchoProxy() (string, bool) {
	proxy := os.Getenv("TGO_DEV_AIR_PROXY_PORT")

	if proxy != "" && proxy != ServerEchoPort {
		return proxy, true
	}

	return "", false
}
