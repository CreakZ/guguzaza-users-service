package cookies

import (
	"guguzaza-users/factory/config"
	"net/http"
	"strconv"
)

type Cooker struct {
	jwtCfg *config.JwtCfg
}

func NewCooker(jwtCfg *config.JwtCfg) Cooker {
	return Cooker{
		jwtCfg: jwtCfg,
	}
}

func (c Cooker) NewJwtCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		MaxAge:   c.jwtCfg.Expiration,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

func (c Cooker) NewIDCookie(id int) *http.Cookie {
	return &http.Cookie{
		Name:     "id",
		Value:    strconv.Itoa(id),
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	}
}
