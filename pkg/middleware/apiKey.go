package middleware

import (
	"net/http"
	"os"
	"strings"
	"health-tech/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (m *middleware) APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" {
			apiKey = c.Query("api_key")
		}

		if apiKey == "" {
			utils.ResponseError(c, http.StatusUnauthorized, "API key diperlukan", nil)
			c.Abort()
			return
		}

		if strings.TrimSpace(apiKey) != os.Getenv("API_KEY") {
			utils.ResponseError(c, http.StatusForbidden, "API key tidak valid", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
