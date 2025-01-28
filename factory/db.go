package factory

import (
	"database/sql"
	"fmt"
	"guguzaza-users/factory/config"
)

func NewDB(cfg *config.PsqlCfg) *sql.DB {
	sslmode := "disable"
	if cfg.EnableSSL {
		sslmode = "enable"
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.DBname,
			sslmode,
		),
	)
	if err != nil {
		panic(fmt.Sprintf("ошибка при подключении к базе данных: %s", err.Error()))
	}

	return db
}
