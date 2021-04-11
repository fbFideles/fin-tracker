package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if authHeader := c.GetHeader("Authorization"); authHeader != "" {
			authHeaderParts := strings.Split(authHeader, " ")
			if len(authHeaderParts) != 2 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "malformed token",
				})
			}

			if authHeaderParts[0] != "Bearer" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Bearer flag is missing, invalid token",
				})
			}

			token, err := jwt.Parse(authHeaderParts[1], func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("token signature method invalid")
				}

				return []byte(os.Getenv("SECRET")), nil
			})

			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
			}

			if !token.Valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "token has been expired",
				})
			}

			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token not provided",
			})
		}
	}
}
