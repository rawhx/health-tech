package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	MoodRepository IMoodRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepoSitory(db),
		MoodRepository: NewMoodRepository(db),
	}
}