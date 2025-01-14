package handlers

import "github.com/labstack/echo/v4"

type MembersHandlers interface {
	RegisterMember(c echo.Context) error
}
