package env

import (
	"os"
	"strings"
)

var (
	BotTelegramToken string
)

func Bot() {
	BotTelegramToken = os.Getenv("TGO_BOT_TELEGRAM_TOKEN")
}

func HasBotTelegramWhitelistUsernames() ([]string, bool) {
	whitelist := os.Getenv("TGO_BOT_TELEGRAM_WHITELIST_USERNAMES")

	if whitelist != "" {
		return strings.Split(whitelist, ","), true
	}

	return nil, false
}
