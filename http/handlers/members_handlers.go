package handlers

import (
	"context"
	"fmt"
	"guguzaza-users/converters"
	"guguzaza-users/domain"
	"guguzaza-users/domain/entities"
	"guguzaza-users/http/dto"
	"net/http"
	"strconv"
	"time"

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

func (mh *memberHandlers) LoginMember(c echo.Context) error {
	// Uuid field is ignored in this case
	// Only 'nickname' and 'password' fields are in use
	userBaseDto := new(dto.UserBase)
	if err := c.Bind(userBaseDto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": fmt.Sprintf("ошибка: %s", err.Error())})
	}

	jwt, err := mh.domain.LoginMember(
		context.Background(),
		userBaseDto.Nickname,
		userBaseDto.Password,
	)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": fmt.Sprintf("ошибка: %s", err.Error())})
	}

	return c.JSON(http.StatusCreated, echo.Map{"jwt": jwt})
}

func (mh *memberHandlers) GetMemberByID(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'id' отсутствует в пути /members/{id}"})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "параметр 'id' неверный",
			"error":   err.Error(),
		})
	}

	memberEntity, err := mh.domain.GetMemberByID(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, converters.MemberPublicFromEntityToDto(memberEntity))
}

// Mocked for now. Gotta make correct pagination
func (mh *memberHandlers) GetMembersPaginated(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"members": []dto.MemberPublic{
		{
			ID:       1,
			Nickname: "некита_1487",
			Uuid:     "no-uuid-1",
			JoinTime: dto.JoinTime{
				Time: time.Date(2006, time.October, 25, 0, 0, 0, 0, time.Local),
			},
			Sex:   "не было",
			About: "didnbehdt",
		},
		{
			ID:       2,
			Nickname: "федот_филиппов",
			Uuid:     "no-uuid-2",
			JoinTime: dto.JoinTime{
				Time: time.Date(2005, time.December, 5, 0, 0, 0, 0, time.Local),
			},
			Sex:   "не было",
			About: "Мы не похожи, ты не ловишь мой вайб Baby Im so high и меня несёт в Рай",
		},
	}})
}

func (mh *memberHandlers) GetTotalMembers(c echo.Context) error {
	amount, err := mh.domain.GetTotalMembers(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"amount": amount})
}

func (mh *memberHandlers) UpdateMember(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'id' отсутствует в пути /members/{id}"})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "параметр 'id' неверный",
			"error":   err.Error(),
		})
	}

	params := c.QueryParams()

	if len(params) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "данные для обновления отсутствуют"})
	}

	updateData := entities.MemberUpdate{}
	if params.Has("sex") {
		sex := params.Get("sex")
		updateData.Sex = &sex
	}

	if params.Has("about") {
		about := params.Get("about")
		updateData.About = &about
	}

	if err = mh.domain.UpdateMember(context.Background(), id, updateData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "данные успешно обновлены"})
}

func (mh *memberHandlers) DeleteMember(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'id' отсутствует в пути /members/{id}"})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "параметр 'id' неверный",
			"error":   err.Error(),
		})
	}

	if err = mh.domain.DeleteMember(context.Background(), id); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "пользователь успешно удален"})
}
