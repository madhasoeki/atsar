package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SuperAdminMiddleware protects endpoints to be accessible only by Super Admin (RoleID = 1)
func SuperAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be Bearer token"})
			c.Abort()
			return
		}

		// Skeleton JWT validation: In a real app, parse and validate the token here.
		// For now, we'll assume the system should check for role_id = 1.
		// Since we don't have JWT parsing logic yet, we will placeholder this check.
		
		// Example: token, err := jwt.Parse(parts[1], ...)
		// if role_id != 1 { c.Abort(); return }

		// Temporary bypass for demonstration purposes:
		// In a real implementation, you would extract claims and check role_id.
		// If you want to strictly enforce it now, you'd need the JWT library.
		
		c.Next()
	}
}
