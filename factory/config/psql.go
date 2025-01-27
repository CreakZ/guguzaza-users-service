package config

type PsqlCfg struct {
	Host      string `toml:"host"`
	Port      int    `toml:"port"`
	User      string `toml:"user"`
	Password  string `toml:"password"`
	DBname    string `toml:"dbname"`
	EnableSSL bool   `toml:"enable_ssl"`
}
