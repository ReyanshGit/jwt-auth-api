package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token lo Header se
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Token nahi hai!"})
			c.Abort()
			return
		}

		// "Bearer TOKEN" se sirf TOKEN nikalo
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Token verify karo
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Token galat hai!"})
			c.Abort()
			return
		}

		// Token se data nikalo
		claims := token.Claims.(jwt.MapClaims)
		c.Set("email", claims["email"])
		c.Set("id", claims["id"])

		c.Next()
	}
}
