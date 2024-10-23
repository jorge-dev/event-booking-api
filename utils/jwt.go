package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates a new JWT token
func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
