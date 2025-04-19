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

var roleHierarchy = map[string]int{
	"USER":      1,
	"MODERATOR": 2,
	"ADMIN":     3,
}

// RequireRoleOrHigher — проверка, что у пользователя роль >= требуемой
func RequireRoleOrHigher(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleAny, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
			c.Abort()
			return
		}

		userRole, ok := roleAny.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid role format"})
			c.Abort()
			return
		}

		userLevel := roleHierarchy[userRole]
		requiredLevel := roleHierarchy[requiredRole]

		if userLevel < requiredLevel {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Удобные обёртки
func RequireUser() gin.HandlerFunc {
	return RequireRoleOrHigher("USER")
}

func RequireModerator() gin.HandlerFunc {
	return RequireRoleOrHigher("MODERATOR")
}

func RequireAdmin() gin.HandlerFunc {
	return RequireRoleOrHigher("ADMIN")
}
