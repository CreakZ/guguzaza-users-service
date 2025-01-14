package main

import (
	"database/sql"
	"fmt"
	"guguzaza-users/http/routing"

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

	routing.InitRouting(e, db)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	if err := e.Start(":8080"); err != nil {
		panic(fmt.Errorf("error while starting the server: %w", err))
	}
}
