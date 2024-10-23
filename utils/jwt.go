package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates a new JWT token
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email":      email,
		"userId":     userId,
		"exp":        time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		errorMessage := fmt.Sprintf("Error generating token: %v", err)
		return "", errors.New(errorMessage)
	}
	return tokenString, nil
}
