package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "авторизационный токен отсутствует"})
		}

		jwt := cookie.Value

		uuid, err := m.jwtUtil.ParseJwtClaims(context.Background(), jwt)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
		}

		c.Set("uuid", uuid)

		return next(c)
	}
}
