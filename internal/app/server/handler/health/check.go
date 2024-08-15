package health

import (
	"net/http"

	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/labstack/echo/v4"
)

func Check(c echo.Context) error {
	if err := c.NoContent(http.StatusOK); err != nil {
		return errs.Response(err, "Check")
	}

	return nil
}
