package routing

import (
	"database/sql"
	token_ports "guguzaza-users/ports/tokens"

	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo, db *sql.DB, jwtUtil token_ports.JwtUtilPort) {
	memberGroup := e.Group("/members")

	InitMembersRouting(memberGroup, db, jwtUtil)
}
