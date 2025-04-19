package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

// AuthMiddleware проверяет JWT-токен
func AuthMiddleware(jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
			c.Abort()
			return
		}

		//TODO fields adapt
		userID, okID := (*claims)["user_id"].(float64)
		userRole, okROLE := (*claims)["user_role"].(string)
		if !okID || !okROLE {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
			c.Abort()
			return
		}
		c.Set("user_role", string(userRole))
		c.Set("user_id", int(userID))
		c.Next()
	}
}

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists || role != "ADMIN" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. ADMIN role required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// RequireModerator проверяет, что роль пользователя — MODERATOR
func RequireModerator() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists || role != "MODERATOR" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. MODERATOR role required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
