package routing

import (
	"database/sql"
	"guguzaza-users/adapters/repository"
	"guguzaza-users/domain"
	"guguzaza-users/http/cookies"
	"guguzaza-users/http/handlers"
	"guguzaza-users/http/middleware"
	ports "guguzaza-users/ports/tokens"

	"github.com/labstack/echo/v4"
)

func InitGeneralRouting(
	e *echo.Echo,
	db *sql.DB,
	middleware middleware.Middleware,
	jwtUtil ports.JwtUtilPort,
	tokensUtil ports.InviteTokensUtilPort,
	cooker cookies.Cooker,
) {
	membersRepo := repository.NewMembersRepository(db)
	adminsRepo := repository.NewAdminsRepository(db)

	membersDomain := domain.NewMembersDomain(membersRepo, jwtUtil)
	adminsDomain := domain.NewAdminsDomain(adminsRepo, jwtUtil, tokensUtil)

	generalHandlers := handlers.NewGeneralHandlers(adminsDomain, membersDomain, cooker)

	loginGroup := e.Group("/login")

	loginGroup.POST("/members", generalHandlers.LoginMember, middleware.ContentTypeMiddleware)
	loginGroup.POST("/admins", generalHandlers.LoginAdmin, middleware.ContentTypeMiddleware)

	e.GET("/stats", generalHandlers.Stats)
}
