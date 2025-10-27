package services

import (
	"errors"
	"health-tech/internal/dto"
	"health-tech/internal/repository"
	"health-tech/models"
	"health-tech/pkg/jwt"
	"health-tech/pkg/utils"

	"github.com/google/uuid"
)

type IUserService interface {
	CreateUser(dto.UserCreateRequest) error
	Login(dto.UserLoginRequest) (dto.LoginResponse, error)
	GetUser(dto.UserParams) (dto.DataUser, error)
}

type UserService struct {
	UserRepository repository.IUserRepository
	JwtAuth        jwt.Interface
}

func NewUserService(userRepository repository.IUserRepository, jwtAuth jwt.Interface) IUserService {
	return &UserService{
		userRepository,
		jwtAuth,
	}
}

func (s *UserService) CreateUser(request dto.UserCreateRequest) error {
	existing, err := s.UserRepository.GetByEmail(request.Email)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("email sudah terdaftar")
	}

	hashedPwd, ok := utils.HashPassword(request.Password)
	if !ok {
		return errors.New("gagal melakukan hash password")
	}

	user := &models.User{
		UserID:   uuid.New().String(),
		Nama:     request.Nama,
		Email:    request.Email,
		Password: hashedPwd,
	}

	return s.UserRepository.CreateUser(user)
}

func (s *UserService) Login(request dto.UserLoginRequest) (dto.LoginResponse, error) {
	var result dto.LoginResponse

	existing, err := s.UserRepository.GetByEmail(request.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}
	if existing == nil {
		return result, errors.New("data tidak ditemukan")
	}

	ok := utils.CheckPassword(request.Password, existing.Password)
	if !ok {
		return dto.LoginResponse{}, errors.New("gagal cek password")
	}

	token, err := s.JwtAuth.CreateJWTToken(existing.UserID)
	if err != nil {
		return result, errors.New("gagal membuat token")
	}

	result.Token = token

	return dto.LoginResponse{
		User: dto.DataUser{
			UserID: existing.UserID,
			Nama:   existing.Nama,
			Email:  existing.Email,
		},
		Token: token,
	}, nil
}

func (s *UserService) GetUser(request dto.UserParams) (dto.DataUser, error) {
	existing, err := s.UserRepository.GetUser(request)
	if err != nil {
		return dto.DataUser{}, err
	}
	if existing == nil {
		return dto.DataUser{}, errors.New("data tidak ditemukan")
	}

	return dto.DataUser{
		UserID: existing.UserID,
		Email:  existing.Email,
		Nama:   existing.Nama,
	}, nil
}
