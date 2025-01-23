package handlers

import "github.com/labstack/echo/v4"

type MembersHandlers interface {
	RegisterMember(c echo.Context) error

	LoginMember(c echo.Context) error

	GetMemberByID(c echo.Context) error
	GetMembersPaginated(c echo.Context) error
	GetTotalMembers(c echo.Context) error

	UpdateMember(c echo.Context) error

	DeleteMember(c echo.Context) error
}

type AdminsHandlers interface {
	CreateInviteToken(c echo.Context) error

	RegisterAdmin(c echo.Context) error

	LoginAdmin(c echo.Context) error

	GetAdminByID(c echo.Context) error
	GetAdminByUuid(c echo.Context) error
	GetAdminsPaginated(c echo.Context) error

	DeleteAdmin(c echo.Context) error
}
