package portfolio

import (
	"net/http"

	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/bastean/tgo/internal/app/server/util/reply"
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/user"
	"github.com/bastean/tgo/internal/pkg/service/user/portfolio"
	"github.com/labstack/echo/v4"
)

func Review(c echo.Context) error {
	primitive := &struct {
		Username string
	}{}

	err := c.Bind(primitive)

	if err != nil {
		return errs.BindingJSON(err, "Review")
	}

	found, err := user.Read.Run(primitive.Username)

	if err != nil {
		return errors.BubbleUp(err, "Review")
	}

	prices, err := portfolio.Price.Run(primitive.Username)

	if err != nil {
		return errors.BubbleUp(err, "Review")
	}

	err = c.JSON(http.StatusOK, &reply.JSON{
		Success: true,
		Message: "Result",
		Data: reply.Payload{
			"Currency": found.Portfolio.Currency,
			"Prices":   prices,
		},
	})

	if err != nil {
		return errs.Response(err, "Review")
	}

	return nil
}
