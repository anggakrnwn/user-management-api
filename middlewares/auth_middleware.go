package middlewares

import (
	"net/http"
	"strings"

	"github.com/anggakrnwn/user-auth-api/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtkey, nil

		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})

			c.Abort()
			return
		}

		c.Set("username", claims.Subject)

		c.Next()
	}
}
