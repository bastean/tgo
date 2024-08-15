package middleware

import (
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

var RateLimiter = middleware.RateLimiter(
	middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	),
)
