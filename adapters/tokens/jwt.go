package tokens

import (
	"context"
	"errors"
	ports "guguzaza-users/ports/tokens"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtUtil struct {
	exp time.Duration
	key []byte
}

func NewJwtUtil(exp time.Duration, key []byte) ports.JwtUtilPort {
	return jwtUtil{
		exp: exp,
		key: key,
	}
}

func (ju jwtUtil) CreateJwt(c context.Context, userUuid string) (token string, err error) {
	expiresAt := time.Now().Add(ju.exp)

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userUuid,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	})

	return jwt.SignedString(ju.key)
}

func (ju jwtUtil) ParseJwtClaims(c context.Context, token string) (userUuid string, err error) {
	claims := new(jwt.RegisteredClaims)

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return ju.key, nil
	})
	if err != nil {
		return "", err
	}

	if !parsedToken.Valid {
		return "", errors.New("срок действия токена истек")
	}

	return claims.Subject, nil
}
