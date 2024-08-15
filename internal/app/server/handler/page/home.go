package page

import (
	"context"

	"github.com/bastean/tgo/internal/app/server/component/page/home"
	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	if err := home.Page().Render(context.Background(), c.Response().Writer); err != nil {
		return errs.Render(err, "Home")
	}

	return nil
}
