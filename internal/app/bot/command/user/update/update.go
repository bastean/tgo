package update

import (
	"github.com/bastean/tgo/internal/app/bot/handler"
	"github.com/bastean/tgo/internal/app/bot/util/errs"
	"github.com/bastean/tgo/internal/pkg/service/user"
	"github.com/bastean/tgo/internal/pkg/service/user/portfolio"
	tele "gopkg.in/telebot.v3"
)

const (
	TotalArgsRequired = 3
)

var Command = tele.Command{
	Text:        "/userupdate",
	Description: "<username> <currency> <coin_list>",
}

func Run(c tele.Context) error {
	args := c.Args()

	err := errs.HasMissingArgs(len(args), TotalArgsRequired, "Run")

	if err != nil {
		return handler.Error(c, err)
	}

	primitive := &user.Primitive{
		Username: args[0],
		Portfolio: &portfolio.Primitive{
			Currency: args[1],
			Coins:    args[2:],
		},
	}

	err = user.Update.Run(primitive)

	if err != nil {
		return handler.Error(c, err)
	}

	err = c.Send("Account updated")

	if err != nil {
		return handler.Error(c, errs.Response(err, "Run"))
	}

	return nil
}
