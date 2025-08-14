package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId string, userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userId,
		"user_name": userName,
		"expiry":    time.Now().Add(365 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte("vsense"))
	if err != nil {
		return "", fmt.Errorf("unable to create auth token")
	}
	return tokenString, err
}
