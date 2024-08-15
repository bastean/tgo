package redirect

import (
	"net/http"

	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/labstack/echo/v4"
)

func Default(c echo.Context) error {
	if err := c.Redirect(http.StatusFound, "/"); err != nil {
		return errs.Response(err, "Default")
	}

	return nil
}
