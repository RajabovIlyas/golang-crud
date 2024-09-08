package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken[T any](payload T, secretJWTKey string, ttl *time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = payload
	if ttl != nil {
		claims["exp"] = now.Add(*ttl).Unix()
	}
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken[T any](tokenString string, secretJWTKey string) (payload T, err error) {
	token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return []byte(secretJWTKey), nil
	})
	if err != nil {
		err = fmt.Errorf("invalid token %v", err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = fmt.Errorf("invalid token")
		return
	}

	value := claims["sub"]

	if value != nil {
		payload = value.(T)
	}

	return
}
