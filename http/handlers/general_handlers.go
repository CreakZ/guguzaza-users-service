package handlers

import (
	"context"
	"fmt"
	"guguzaza-users/domain"
	"guguzaza-users/http/cookies"
	"guguzaza-users/http/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type generalHandlers struct {
	adminsDomain  domain.AdminsDomain
	membersDomain domain.MembersDomain
}

func NewGeneralHandlers(
	adminsDomain domain.AdminsDomain,
	membersDomain domain.MembersDomain,
) GeneralHandlers {
	return generalHandlers{
		adminsDomain:  adminsDomain,
		membersDomain: membersDomain,
	}
}

func (gh generalHandlers) LoginAdmin(c echo.Context) error {
	creds := new(dto.Credentials)

	if err := c.Bind(creds); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": fmt.Sprintf("что-то пошло не так: %s", err.Error()),
		})
	}

	jwt, err := gh.adminsDomain.LoginAdmin(context.Background(), creds.Nickname, creds.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	c.SetCookie(cookies.NewJwtCookie(jwt))

	return c.JSON(http.StatusCreated, echo.Map{"message": "успешно"})
}

func (gh generalHandlers) LoginMember(c echo.Context) error {
	creds := new(dto.Credentials)
	if err := c.Bind(creds); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": fmt.Sprintf("ошибка: %s", err.Error())})
	}

	jwt, err := gh.membersDomain.LoginMember(
		context.Background(),
		creds.Nickname,
		creds.Password,
	)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
	}

	c.SetCookie(cookies.NewJwtCookie(jwt))

	return c.JSON(http.StatusCreated, echo.Map{"message": "успешно"})
}

func (gh generalHandlers) Stats(c echo.Context) error {
	return c.NoContent(http.StatusServiceUnavailable)
}
