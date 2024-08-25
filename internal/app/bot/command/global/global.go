package global

import (
	"github.com/bastean/tgo/internal/app/bot/command/global/help"
	"github.com/bastean/tgo/internal/app/bot/command/global/start"
	"github.com/bastean/tgo/internal/app/bot/handler"
)

var Routing = handler.Command{
	start.Command: help.Run,
	help.Command:  help.Run,
}
