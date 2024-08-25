package review

import (
	"fmt"
	"strings"

	"github.com/bastean/tgo/internal/app/bot/handler"
	"github.com/bastean/tgo/internal/app/bot/util/errs"
	"github.com/bastean/tgo/internal/pkg/service/user/portfolio"
	tele "gopkg.in/telebot.v3"
)

const (
	TotalArgsRequired = 1
)

var Command = tele.Command{
	Text:        "/portfolioreview",
	Description: "<username>",
}

func Run(c tele.Context) error {
	args := c.Args()

	err := errs.HasMissingArgs(len(args), TotalArgsRequired, "Run")

	if err != nil {
		return handler.Error(c, err)
	}

	username := args[0]

	prices, err := portfolio.Price.Run(username)

	if err != nil {
		return handler.Error(c, err)
	}

	result := `*Result*
	
*Prices:*
`

	for coin, price := range prices {
		result += fmt.Sprintf(`
\- *%s:* %g
`, coin, price)
	}

	result = strings.ReplaceAll(result, ".", "\\.")

	err = c.Send(result, &tele.SendOptions{
		ParseMode: "MarkdownV2",
	})

	if err != nil {
		return handler.Error(c, errs.Response(err, "Run"))
	}

	return nil
}
