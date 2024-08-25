package user

import (
	"github.com/bastean/tgo/internal/app/bot/command/user/create"
	"github.com/bastean/tgo/internal/app/bot/command/user/delete"
	"github.com/bastean/tgo/internal/app/bot/command/user/read"
	"github.com/bastean/tgo/internal/app/bot/command/user/update"
	"github.com/bastean/tgo/internal/app/bot/handler"
)

var Routing = handler.Command{
	create.Command: create.Run,
	read.Command:   read.Run,
	update.Command: update.Run,
	delete.Command: delete.Run,
}
