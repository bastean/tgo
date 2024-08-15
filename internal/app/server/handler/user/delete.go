package user

import (
	"net/http"

	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/bastean/tgo/internal/app/server/util/reply"
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/user"
	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	primitive := &struct {
		Username string
	}{}

	err := c.Bind(primitive)

	if err != nil {
		return errs.BindingJSON(err, "Delete")
	}

	err = user.Delete.Run(primitive.Username)

	if err != nil {
		return errors.BubbleUp(err, "Delete")
	}

	err = c.JSON(http.StatusOK, &reply.JSON{
		Success: true,
		Message: "Account deleted",
	})

	if err != nil {
		return errs.Response(err, "Delete")
	}

	return nil
}
