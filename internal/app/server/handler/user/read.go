package user

import (
	"net/http"

	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/bastean/tgo/internal/app/server/util/reply"
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/user"
	"github.com/labstack/echo/v4"
)

func Read(c echo.Context) error {
	primitive := &struct {
		Username string
	}{}

	err := c.Bind(primitive)

	if err != nil {
		return errs.BindingJSON(err, "Read")
	}

	found, err := user.Read.Run(primitive.Username)

	if err != nil {
		return errors.BubbleUp(err, "Read")
	}

	err = c.JSON(http.StatusOK, &reply.JSON{
		Success: true,
		Message: "Found",
		Data: reply.Payload{
			"Username":  found.Username,
			"Portfolio": found.Portfolio,
		},
	})

	if err != nil {
		return errs.Response(err, "Read")
	}

	return nil
}
