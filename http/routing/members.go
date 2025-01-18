package routing

import (
	"database/sql"
	"guguzaza-users/adapters/repository"
	"guguzaza-users/domain"
	"guguzaza-users/http/handlers"
	token_ports "guguzaza-users/ports/tokens"

	"github.com/labstack/echo/v4"
)

func InitMembersRouting(e *echo.Group, db *sql.DB, jwtUtil token_ports.JwtUtilPort) {
	membersRepo := repository.NewMembersRepository(db)
	membersDomain := domain.NewMembersDomain(membersRepo, jwtUtil)
	membersHandlers := handlers.NewMembersHandlers(membersDomain)

	e.POST("", membersHandlers.RegisterMember)

	e.POST("/login", membersHandlers.LoginMember)

	e.GET("/:id", membersHandlers.GetMemberByID)
	e.GET("", membersHandlers.GetMembersPaginated)
	e.GET("/amount", membersHandlers.GetTotalMembers)

	e.PATCH("/:id", membersHandlers.UpdateMember)

	e.DELETE("/:id", membersHandlers.DeleteMember)
}
