package router

import (
	"github.com/bastean/tgo/internal/app/bot/command"
	tele "gopkg.in/telebot.v3"
)

func Updates(session *tele.Bot) {
	for _, routes := range command.Routing {
		for command, handler := range routes {
			session.Handle(command.Text, handler)
		}
	}
}
