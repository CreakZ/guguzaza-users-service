package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hash string, err error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashBytes), nil
}
