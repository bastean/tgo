package read

import (
	"fmt"

	"github.com/bastean/tgo/internal/app/bot/handler"
	"github.com/bastean/tgo/internal/app/bot/util/errs"
	"github.com/bastean/tgo/internal/app/bot/util/escape"
	"github.com/bastean/tgo/internal/pkg/service/user"
	tele "gopkg.in/telebot.v3"
)

const (
	TotalArgsRequired = 1
)

var Command = tele.Command{
	Text:        "/userread",
	Description: "<username>",
}

func Run(c tele.Context) error {
	args := c.Args()

	err := errs.HasMissingArgs(len(args), TotalArgsRequired, "Run")

	if err != nil {
		return handler.Error(c, err)
	}

	username := args[0]

	found, err := user.Read.Run(username)

	if err != nil {
		return handler.Error(c, err)
	}

	result := fmt.Sprintf(`*Found*

*Username:* %s

*Portfolio*

*Currency:* %s

*Coins:* %s
`, found.Username, found.Portfolio.Currency, found.Portfolio.Coins)

	result = escape.ReservedCharacters(result)

	err = c.Send(result, &tele.SendOptions{ParseMode: "MarkdownV2"})

	if err != nil {
		return handler.Error(c, errs.Response(err, "Run"))
	}

	return nil
}
