package middleware

import (
	"guguzaza-users/factory/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (m Middleware) CorsMiddleware(cfg *config.FrontendCfg) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: cfg.Origins,
		},
	)
}
