package router

import (
	"embed"

	"github.com/bastean/tgo/internal/app/server/handler/health"
	"github.com/bastean/tgo/internal/app/server/handler/redirect"
	"github.com/bastean/tgo/internal/app/server/middleware"
	"github.com/bastean/tgo/internal/app/server/router/api"
	"github.com/bastean/tgo/internal/app/server/router/view"
	"github.com/labstack/echo/v4"
)

var Router *echo.Echo

func New(files *embed.FS) *echo.Echo {
	Router = echo.New()

	Router.HTTPErrorHandler = middleware.Error

	Router.Use(middleware.Recover)

	Router.Use(middleware.Headers)

	Router.Use(middleware.RateLimiter)

	Router.StaticFS("/public", files)

	api.Use(Router.Group("/v0"))

	view.Use(Router)

	Router.HEAD("/health", health.Check)

	Router.RouteNotFound("/*", redirect.Default)

	return Router
}
