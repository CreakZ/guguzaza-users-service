package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m Middleware) ContentTypeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headers := c.Request().Header

		contentType := headers.Get("Content-Type")
		switch contentType {
		case "":
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "заголовок Content-Type отсутствует"})

		case echo.MIMEApplicationJSON, echo.MIMEMultipartForm:
			return next(c)

		default:
			return c.JSON(http.StatusUnprocessableEntity, echo.Map{
				"message": fmt.Sprintf("заголовок Content-Type неверен: %s", contentType),
			})
		}
	}
}
