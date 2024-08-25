package help

import (
	tele "gopkg.in/telebot.v3"
)

var Command = tele.Command{
	Text:        "/help",
	Description: "Show help",
}

func Run(c tele.Context) error {
	return c.Send(`*Example of interoperability between a Web App, a Telegram Bot and a third\-party API\.*

*Use the "Menu" button or "/" to display the available commands\.*

[Powered by CoinGecko API](https://www.coingecko.com/api)`,
		&tele.SendOptions{
			ParseMode: "MarkdownV2",
		})
}
