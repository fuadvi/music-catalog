package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(10 * time.Minute).Unix(),
	})

	key := []byte(secretKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := make(jwt.MapClaims)

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", err
	}

	return int64(claims["id"].(float64)), claims["username"].(string), nil
}

func ValidateTokenWithoutExpiry(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := make(jwt.MapClaims)

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithoutClaimsValidation())

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", err
	}

	return int64(claims["id"].(float64)), claims["username"].(string), nil
}
