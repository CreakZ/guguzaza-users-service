package routing

import (
	"database/sql"
	"guguzaza-users/adapters/repository"
	"guguzaza-users/domain"
	"guguzaza-users/http/handlers"
	token_ports "guguzaza-users/ports/tokens"

	"github.com/labstack/echo/v4"
)

func InitAdminsRouting(
	e *echo.Group,
	db *sql.DB,
	jwtUtil token_ports.JwtUtilPort,
	tokensUtil token_ports.InviteTokensUtilPort,
) {
	adminsRepo := repository.NewAdminsRepository(db)
	adminsDomain := domain.NewAdminsDomain(adminsRepo, jwtUtil, tokensUtil)
	adminsHandlers := handlers.NewAdminsHandlers(adminsDomain)

	e.POST("", adminsHandlers.RegisterAdmin)
	e.POST("/login", adminsHandlers.LoginAdmin)
	e.GET("/:id", adminsHandlers.GetAdminByID)
	e.GET("/me", adminsHandlers.GetAdminByUuid)
	e.GET("", adminsHandlers.GetAdminsPaginated)
	e.DELETE("", adminsHandlers.DeleteAdmin)
}
