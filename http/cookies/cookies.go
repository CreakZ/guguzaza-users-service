package cookies

import "net/http"

func NewJwtCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "jwt",
		Value:    token,
		MaxAge:   60 * 60 * 24 * 7, // 7 days (60 * 60 * 24 * 7 seconds = 7 days)
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}
