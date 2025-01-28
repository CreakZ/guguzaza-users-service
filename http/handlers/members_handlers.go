package handlers

import (
	"context"
	"guguzaza-users/converters"
	"guguzaza-users/domain"
	"guguzaza-users/domain/entities"
	"guguzaza-users/http/cookies"
	"guguzaza-users/http/dto"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type memberHandlers struct {
	domain domain.MembersDomain
	cooker cookies.Cooker
}

func NewMembersHandlers(domain domain.MembersDomain, cooker cookies.Cooker) MembersHandlers {
	return &memberHandlers{
		domain: domain,
		cooker: cooker,
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

	c.SetCookie(mh.cooker.NewIDCookie(id))

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "пользователь создан успешно",
	})
}

func (mh *memberHandlers) GetMemberByID(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'id' отсутствует в пути /members/{id}"})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "параметр 'id' приведен в неверном формате",
		})
	}

	memberEntity, err := mh.domain.GetMemberByID(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, converters.MemberPublicFromEntityToDto(memberEntity))
}

func (mh *memberHandlers) GetMembersPaginated(c echo.Context) error {
	queryParams := c.QueryParams()

	if !queryParams.Has("page") {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'page' отсутствует"})
	}

	if !queryParams.Has("limit") {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'limit' отсутствует"})
	}

	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'page' приведен в некорректном формате"})
	}

	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'limit' приведен в некорректном формате"})
	}

	members, err := mh.domain.GetMembersPaginated(context.Background(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"limit":   limit,
		"page":    page,
		"members": members,
	})
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
			"message": "параметр 'id' приведен в неверном формате",
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
