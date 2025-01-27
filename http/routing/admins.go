package routing

import (
	"database/sql"
	"guguzaza-users/adapters/repository"
	"guguzaza-users/domain"
	"guguzaza-users/http/handlers"
	"guguzaza-users/http/middleware"
	token_ports "guguzaza-users/ports/tokens"

	"github.com/labstack/echo/v4"
)

func InitAdminsRouting(
	e *echo.Group,
	db *sql.DB,
	middleware middleware.Middleware,
	jwtUtil token_ports.JwtUtilPort,
	tokensUtil token_ports.InviteTokensUtilPort,
) {
	adminsRepo := repository.NewAdminsRepository(db)
	adminsDomain := domain.NewAdminsDomain(adminsRepo, jwtUtil, tokensUtil)
	adminsHandlers := handlers.NewAdminsHandlers(adminsDomain)

	e.POST("", adminsHandlers.RegisterAdmin)
	e.GET("/:id", adminsHandlers.GetAdminByID)
	e.GET("/me", adminsHandlers.GetAdminByUuid, middleware.JwtMiddleware)
	e.GET("", adminsHandlers.GetAdminsPaginated)
	e.DELETE("", adminsHandlers.DeleteAdmin)
}
