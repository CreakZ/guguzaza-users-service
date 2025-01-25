package handlers

import (
	"context"
	"fmt"
	"guguzaza-users/converters"
	"guguzaza-users/domain"
	"guguzaza-users/http/dto"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type adminsHandlers struct {
	adminsDomain domain.AdminsDomain
}

func NewAdminsHandlers(adminsDomain domain.AdminsDomain) AdminsHandlers {
	return adminsHandlers{
		adminsDomain: adminsDomain,
	}
}

func (ah adminsHandlers) CreateInviteToken(c echo.Context) error {
	return c.JSON(http.StatusServiceUnavailable, echo.Map{"message": "пока что это не работает"}) // гугузаза токены
}

func (ah adminsHandlers) RegisterAdmin(c echo.Context) error {
	admin := new(dto.AdminCreate)
	if err := c.Bind(admin); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": fmt.Sprintf("что-то пошло не так: %s", err.Error())})
	}

	id, err := ah.adminsDomain.RegisterAdmin(
		context.Background(),
		converters.AdminCreateFromDtoToEntity(*admin),
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	c.SetCookie(&http.Cookie{
		Name:     "id",
		Value:    strconv.Itoa(id),
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	})

	return c.JSON(http.StatusCreated, echo.Map{"message": "администратор создан успешно"})
}

func (ah adminsHandlers) GetAdminByID(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'id' отсутствует в пути /admins/{id}"})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "параметр 'id' приведен в неверном формате",
		})
	}

	admin, err := ah.adminsDomain.GetAdminByID(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, converters.AdminPublicFromEntityToDto(admin))
}

func (ah adminsHandlers) GetAdminByUuid(c echo.Context) error {
	uuid := c.Get("uuid").(string)

	admin, err := ah.adminsDomain.GetAdminByUuid(context.Background(), uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, converters.AdminPublicFromEntityToDto(admin))
}

func (ah adminsHandlers) GetAdminsPaginated(c echo.Context) error {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")
	if pageStr == "" || limitStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'page' или 'limit' отсутствует"})
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": fmt.Sprintf("параметр 'page' не является целым числом: %s", pageStr),
		})
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": fmt.Sprintf("параметр 'limit' не является целым числом: %s", limitStr),
		})
	}

	admins, err := ah.adminsDomain.GetAdminsPaginated(
		context.Background(),
		int64(page),
		int64(limit),
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, converters.AdminsPaginatedFromEntityToDto(admins))
}

func (ah adminsHandlers) DeleteAdmin(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "параметр 'id' отсутствует в пути /admins/{id}"})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "параметр 'id' приведен в неверном формате",
		})
	}

	err = ah.adminsDomain.DeleteAdminByID(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "пользователь успешно удален"})
}
