// middleware/auth.go

package middleware

import (
	"net/http"
	"your-package/models"
	"your-package/config"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware to authenticate users using JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userID, err := models.ValidateToken(token, config.SecretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
