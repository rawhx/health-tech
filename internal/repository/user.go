package repository

import (
	"health-tech/internal/dto"
	"health-tech/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(*models.User) error
	GetByEmail(string) (*models.User, error)
	GetUser(dto.UserParams) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepoSitory(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUser(params dto.UserParams) (*models.User, error) {
	var user models.User
	query := r.db.Model(&models.User{})

	if params.UserID != "" {
		query = query.Where("id = ?", params.UserID)
	}
	if params.Email != "" {
		query = query.Where("email = ?", params.Email)
	}
	if params.Nama != "" {
		query = query.Where("nama = ?", params.Nama)
	}

	err := query.First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
