package routing

import (
	"database/sql"
	"guguzaza-users/http/cookies"
	"guguzaza-users/http/middleware"
	token_ports "guguzaza-users/ports/tokens"

	"github.com/labstack/echo/v4"
)

func InitRouting(
	e *echo.Echo,
	db *sql.DB,
	middleware middleware.Middleware,
	jwtUtil token_ports.JwtUtilPort,
	tokensUtil token_ports.InviteTokensUtilPort,
	cooker cookies.Cooker,
) {
	memberGroup := e.Group("/members")
	admGroup := e.Group("/admins")

	InitMembersRouting(memberGroup, db, middleware, jwtUtil, cooker)
	InitAdminsRouting(admGroup, db, middleware, jwtUtil, tokensUtil, cooker)
	InitGeneralRouting(e, db, middleware, jwtUtil, tokensUtil, cooker)
}
