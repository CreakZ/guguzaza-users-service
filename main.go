package main

import (
	"fmt"
	"guguzaza-users/adapters/repository"
	"guguzaza-users/adapters/tokens"
	"guguzaza-users/factory"
	"guguzaza-users/factory/config"
	"guguzaza-users/http/cookies"
	mw "guguzaza-users/http/middleware"
	"guguzaza-users/http/routing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	cfg := config.NewConfig()

	db := factory.NewDB(cfg.Postgres)

	key := factory.NewKey()

	jwtUtil := tokens.NewJwtUtil(time.Second*time.Duration(cfg.Jwt.Expiration), key)
	tokensUtil := tokens.NewInviteTokensUtil(repository.NewInviteTokensRepository(db))

	middleW := mw.NewMiddleware(jwtUtil)

	e.Use(middleware.Logger())
	e.Use(middleW.CorsMiddleware(cfg.Frontend))

	cooker := cookies.NewCooker(cfg.Jwt)

	routing.InitRouting(e, db, middleW, jwtUtil, tokensUtil, cooker)

	if err := e.Start(":8080"); err != nil {
		panic(fmt.Errorf("ошибка при запуске сервера: %w", err))
	}
}
