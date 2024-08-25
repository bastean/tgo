package command

import (
	"github.com/bastean/tgo/internal/app/bot/command/global"
	"github.com/bastean/tgo/internal/app/bot/command/portfolio"
	"github.com/bastean/tgo/internal/app/bot/command/user"
	"github.com/bastean/tgo/internal/app/bot/handler"
	tele "gopkg.in/telebot.v3"
)

var Routing = []handler.Command{
	global.Routing,
	user.Routing,
	portfolio.Routing,
}

func List() []tele.Command {
	commands := []tele.Command{}

	for _, routes := range Routing {
		for command := range routes {
			commands = append(commands, command)
		}
	}

	return commands
}
