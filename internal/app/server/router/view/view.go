package view

import (
	"github.com/bastean/tgo/internal/app/server/handler/page"
	"github.com/labstack/echo/v4"
)

func Use(router *echo.Echo) {
	router.GET("/", page.Home)
}
