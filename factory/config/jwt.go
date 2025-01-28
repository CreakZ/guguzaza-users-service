package config

type JwtCfg struct {
	Expiration int `toml:"expiration"` // expiration is measured in seconds
}
