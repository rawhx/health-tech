package handler

import (
	"fmt"
	"health-tech/internal/services"
	"health-tech/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	services   *services.Service
	middleware middleware.Middleware
}

func NewRest(services *services.Service, middleware middleware.Middleware) *Rest {
	return &Rest{
		router:     gin.Default(),
		services:   services,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	r.router.Use(r.middleware.Cors())

	routerGroup := r.router.Group("api/v1", r.middleware.APIKeyAuthMiddleware())
	auth := routerGroup.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/login", r.Login)

	user := routerGroup.Group("/mood")
	user.POST("/", r.CreateMood)
	user.GET("/:user_id", r.GetUserMoods)
	user.GET("/summary/:user_id", r.GetMoodSummary)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
