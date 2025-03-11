package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, expiry time.Duration, tokenType string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     "staff", 
		"exp":      time.Now().Add(expiry).Unix(),
		"type":     tokenType, 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var secretKey []byte
	if tokenType == "access" {
		secretKey = []byte(os.Getenv("JWT_ACCESS_SECRET"))
	} else {
		secretKey = []byte(os.Getenv("JWT_REFRESH_SECRET"))
	}

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string, tokenType string) (*jwt.MapClaims, error) {
	var secretKey []byte
	if tokenType == "access" {
		secretKey = []byte(os.Getenv("JWT_ACCESS_SECRET"))
	} else {
		secretKey = []byte(os.Getenv("JWT_REFRESH_SECRET"))
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != tokenType {
		return nil, err
	}

	return &claims, nil
}