// Ficheiro: internal/auth/auth.go

package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// O nosso Middleware agora vive no seu próprio pacote.
// (Repare que o nome começa com letra maiúscula para ser exportado)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Header de autorização em falta"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}
			// Agora usamos a variável deste pacote
			return JwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token inválido"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := int(claims["sub"].(float64))
			c.Set("userId", userID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Claims do token inválidas"})
		}
	}
}
