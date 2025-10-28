package middleware

import (
	"health-tech/internal/services"
	"health-tech/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthenticateUser(c *gin.Context)
	APIKeyAuthMiddleware() gin.HandlerFunc
	Cors() gin.HandlerFunc
}

type middleware struct {
	services *services.Service
	jwtAuth jwt.Interface
}

func Init(service *services.Service, jwtAuth jwt.Interface) Middleware {
	return &middleware{
		services: service,
		jwtAuth: jwtAuth,
	}
}
