package user

import (
	"net/http"

	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/bastean/tgo/internal/app/server/util/reply"
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/user"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	primitive := new(user.Primitive)

	err := c.Bind(primitive)

	if err != nil {
		return errs.BindingJSON(err, "Create")
	}

	err = user.Create.Run(primitive)

	if err != nil {
		return errors.BubbleUp(err, "Create")
	}

	err = c.JSON(http.StatusCreated, &reply.JSON{
		Success: true,
		Message: "Account created",
	})

	if err != nil {
		return errs.Response(err, "Create")
	}

	return nil
}
