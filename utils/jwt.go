package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "secret"

// GenerateToken generates a new JWT token
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email":      email,
		"userId":     userId,
		"exp":        time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		errorMessage := fmt.Sprintf("Error generating token: %v", err)
		return "", errors.New(errorMessage)
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		errorMessage := fmt.Sprintf("Error parsing token: %v", err)
		return 0, errors.New(errorMessage)
	}
	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("error getting claims")
	}

	// Convert userId to int64
	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("userId claim is not a valid number")
	}
	userId := int64(userIdFloat)

	return userId, nil

}
