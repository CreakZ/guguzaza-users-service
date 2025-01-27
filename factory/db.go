package factory

import (
	"database/sql"
	"fmt"
	"guguzaza-users/factory/config"
)

func NewDB(cfg *config.Cfg) *sql.DB {
	sslmode := "disable"
	if cfg.Postgres.EnableSSL {
		sslmode = "enable"
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.DBname,
			sslmode,
		),
	)
	if err != nil {
		panic(fmt.Sprintf("ошибка при подключении к базе данных: %s", err.Error()))
	}

	return db
}
