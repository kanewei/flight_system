package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		const apiKey = "your_api_key_here"
		// Check for API key in request headers
		if key := c.GetHeader("API-Key"); key != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
