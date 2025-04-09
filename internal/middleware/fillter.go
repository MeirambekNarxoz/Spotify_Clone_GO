package middleware

import (
	"Spotify/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware проверяет JWT-токен
func AuthMiddleware(jwtService *service.JWTService, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// Проверяем роль, если требуется
		if requiredRole != "" {
			role, ok := claims["authorities"].(string)
			if !ok || role != requiredRole {
				c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
				c.Abort()
				return
			}
		}
		// Извлекаем userID из токена и сохраняем его в контексте запроса
		userID, ok := claims["id"].(float64) // или другой тип в зависимости от структуры твоего токена
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
