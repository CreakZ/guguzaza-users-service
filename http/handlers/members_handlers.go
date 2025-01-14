package handlers

import (
	"context"
	"fmt"
	"guguzaza-users/converters"
	"guguzaza-users/domain"
	"guguzaza-users/http/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type memberHandlers struct {
	domain domain.MembersDomain
}

func NewMembersHandlers(domain domain.MembersDomain) MembersHandlers {
	return &memberHandlers{
		domain: domain,
	}
}

func (mh *memberHandlers) RegisterMember(c echo.Context) error {
	memberDto := new(dto.MemberCreate)
	if err := c.Bind(memberDto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "ошибка в теле запроса",
			"error":   err.Error(),
		})
	}

	fmt.Println(memberDto)

	memberEntity := converters.MemberCreateFromDtoToEntity(*memberDto)

	id, err := mh.domain.RegisterMember(context.Background(), memberEntity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id":      id,
		"message": "пользователь создан успешно",
	})
}
