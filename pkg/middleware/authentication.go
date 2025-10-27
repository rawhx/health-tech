package middleware

import (
	"health-tech/internal/dto"
	"health-tech/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		utils.ResponseError(c, http.StatusUnauthorized, "token kosong", nil)
		c.Abort()
		return
	}

	token := strings.Split(bearer, " ")[1]
	userID, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, "gagal melakukan validasi token", err)
		c.Abort()
		return
	}

	user, err := m.services.UserService.GetUser(dto.UserParams{
		UserID: userID,
	})
	if err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, "gagal mendapatkan data user", err)
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}