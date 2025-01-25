package routing

import (
	"database/sql"
	"guguzaza-users/adapters/repository"
	"guguzaza-users/domain"
	"guguzaza-users/http/handlers"
	ports "guguzaza-users/ports/tokens"

	"github.com/labstack/echo/v4"
)

func InitGeneralRouting(e *echo.Echo, db *sql.DB, jwtUtil ports.JwtUtilPort, tokensUtil ports.InviteTokensUtilPort) {
	membersRepo := repository.NewMembersRepository(db)
	adminsRepo := repository.NewAdminsRepository(db)

	membersDomain := domain.NewMembersDomain(membersRepo, jwtUtil)
	adminsDomain := domain.NewAdminsDomain(adminsRepo, jwtUtil, tokensUtil)

	generalHandlers := handlers.NewGeneralHandlers(adminsDomain, membersDomain)

	loginGroup := e.Group("/login")

	loginGroup.POST("/members", generalHandlers.LoginMember)
	loginGroup.POST("/admins", generalHandlers.LoginAdmin)

	e.GET("/stats", generalHandlers.Stats)
}
