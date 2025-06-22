package helpers

import (
	"time"

	"github.com/anggakrnwn/user-auth-api/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func GenerateToken(username string) string {
	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtkey)

	return token
}
