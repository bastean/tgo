package start

import (
	tele "gopkg.in/telebot.v3"
)

var Command = tele.Command{
	Text:        "/start",
	Description: "Start tGO",
}
