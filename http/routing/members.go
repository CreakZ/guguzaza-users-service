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

func InitMembersRouting(e *echo.Group, db *sql.DB, middleware middleware.Middleware, jwtUtil token_ports.JwtUtilPort) {
	membersRepo := repository.NewMembersRepository(db)
	membersDomain := domain.NewMembersDomain(membersRepo, jwtUtil)
	membersHandlers := handlers.NewMembersHandlers(membersDomain)

	e.POST("", membersHandlers.RegisterMember, middleware.ContentTypeMiddleware)

	e.GET("/:id", membersHandlers.GetMemberByID)
	e.GET("", membersHandlers.GetMembersPaginated)

	e.PATCH("/:id", membersHandlers.UpdateMember, middleware.ContentTypeMiddleware)

	e.DELETE("/:id", membersHandlers.DeleteMember, middleware.ContentTypeMiddleware)
}
