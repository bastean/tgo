package user

import (
	"net/http"

	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/bastean/tgo/internal/app/server/util/reply"
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/user"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	primitive := new(user.Primitive)

	err := c.Bind(primitive)

	if err != nil {
		return errs.BindingJSON(err, "Update")
	}

	err = user.Update.Run(primitive)

	if err != nil {
		return errors.BubbleUp(err, "Update")
	}

	err = c.JSON(http.StatusOK, &reply.JSON{
		Success: true,
		Message: "Account updated",
	})

	if err != nil {
		return errs.Response(err, "Update")
	}

	return nil
}
