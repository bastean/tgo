package delete

import (
	"github.com/bastean/tgo/internal/app/bot/handler"
	"github.com/bastean/tgo/internal/app/bot/util/errs"
	"github.com/bastean/tgo/internal/pkg/service/user"
	tele "gopkg.in/telebot.v3"
)

const (
	TotalArgsRequired = 1
)

var Command = tele.Command{
	Text:        "/userdelete",
	Description: "<username>",
}

func Run(c tele.Context) error {
	args := c.Args()

	err := errs.HasMissingArgs(len(args), TotalArgsRequired, "Run")

	if err != nil {
		return handler.Error(c, err)
	}

	username := args[0]

	err = user.Delete.Run(username)

	if err != nil {
		return handler.Error(c, err)
	}

	err = c.Send("Account deleted")

	if err != nil {
		return handler.Error(c, errs.Response(err, "Run"))
	}

	return nil
}
