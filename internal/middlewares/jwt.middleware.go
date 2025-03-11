package middlewares

import (
	"net/http"
	"stockhive-server/internal/config"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(tokenType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims, err := config.VerifyToken(tokenString, tokenType)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		mapClaims := *claims // dereference pointer jadi map
		if username, ok := mapClaims["username"].(string); ok {
			c.Set("username", username)
		}
		if role, ok := mapClaims["role"].(string); ok {
			c.Set("role", role)
		}
		c.Next()
		
	}
}