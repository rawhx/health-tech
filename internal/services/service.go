package services

import (
	"health-tech/internal/repository"
	"health-tech/pkg/jwt"
)

type Service struct {
	UserService IUserService
	jwtAuth     jwt.Interface
}

func NewService(repository *repository.Repository, jwtAuth jwt.Interface) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository, jwtAuth),
	}
}
