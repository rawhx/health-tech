package services

import (
	"health-tech/internal/repository"
	"health-tech/pkg/jwt"
)

type Service struct {
	UserService IUserService
	MoodService IMoodService
}

func NewService(repository *repository.Repository, jwtAuth jwt.Interface) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository, jwtAuth),
		MoodService: NewMoodService(repository.MoodRepository, repository.UserRepository),
	}
}
