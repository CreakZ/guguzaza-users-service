package routing

import (
	"database/sql"
	token_ports "guguzaza-users/ports/tokens"

	"github.com/labstack/echo/v4"
)

func InitRouting(
	e *echo.Echo,
	db *sql.DB,
	jwtUtil token_ports.JwtUtilPort,
	tokensUtil token_ports.InviteTokensUtilPort,
) {
	memberGroup := e.Group("/members")
	admGroup := e.Group("/admins")

	InitMembersRouting(memberGroup, db, jwtUtil)
	InitAdminsRouting(admGroup, db, jwtUtil, tokensUtil)
}
