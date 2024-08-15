package api

import (
	"github.com/bastean/tgo/internal/app/server/handler/portfolio"
	"github.com/bastean/tgo/internal/app/server/handler/user"
	"github.com/labstack/echo/v4"
)

func Use(router *echo.Group) {
	router.PUT("/user", user.Create)
	router.POST("/user", user.Read)
	router.PATCH("/user", user.Update)
	router.DELETE("/user", user.Delete)

	router.POST("/portfolio", portfolio.Review)
}
