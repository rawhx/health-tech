package main

import (
	"health-tech/internal/handler"
	"health-tech/internal/repository"
	"health-tech/internal/services"
	"health-tech/pkg/config"
	"health-tech/pkg/database"
	"health-tech/pkg/jwt"
	"health-tech/pkg/middleware"
	"log"
)

func main() {
	if err := config.LoadEnvironment(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connect database: %v", err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	repository := repository.NewRepository(db)
	jwt := jwt.Init()
	services := services.NewService(repository, jwt)
	middleware := middleware.Init(services, jwt)

	r := handler.NewRest(services, middleware)
	r.MountEndpoint()
	r.Run()

}