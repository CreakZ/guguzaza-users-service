package config

type Jwt struct {
	Expiration int `toml:"expiration"` // expiration is measured in seconds
}
