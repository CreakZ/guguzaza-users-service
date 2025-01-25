package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"guguzaza-users/adapters/repository"
	"guguzaza-users/adapters/tokens"
	"guguzaza-users/http/routing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	db, err := sql.Open("postgres", "host=localhost port=5433 user=postgres password=guguzaza dbname=guguzaza sslmode=disable")
	if err != nil {
		panic(err)
	}

	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		panic(fmt.Sprintf("Ошибка генерации ключа: %v", err))
	}

	jwtUtil := tokens.NewJwtUtil(time.Second*10, key)
	tokensUtil := tokens.NewInviteTokensUtil(repository.NewInviteTokensRepository(db))

	routing.InitRouting(e, db, jwtUtil, tokensUtil)

	if err := e.Start(":8080"); err != nil {
		panic(fmt.Errorf("error while starting the server: %w", err))
	}
}
