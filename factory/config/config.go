package config

import (
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

type Cfg struct {
	Postgres *PsqlCfg `toml:"postgres"`
	Jwt      *Jwt     `toml:"jwt"`
	// Frontend *FrontendCfg `toml:"frontend"`
}

func NewConfig() (cfg *Cfg) {
	cfgFile, err := os.Open("cfg.toml")
	if err != nil {
		panic(fmt.Sprintf("ошибка при открытии файла: %s", err.Error()))
	}

	data, err := io.ReadAll(cfgFile)
	if err != nil {
		panic(fmt.Sprintf("ошибка при чтении файла: %s", err.Error()))
	}

	cfg = new(Cfg)
	if _, err = toml.Decode(string(data), cfg); err != nil {
		panic(fmt.Sprintf("ошибка при декодировании файла: %s", err.Error()))
	}

	return cfg
}
