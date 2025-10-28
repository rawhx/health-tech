package repository

import (
	"errors"
	"health-tech/models"
	"health-tech/pkg/pagination"
	"time"

	"gorm.io/gorm"
)

type IMoodRepository interface {
	CreateMood(*models.Mood) error
	GetUserMoods(string, pagination.Params) ([]models.Mood, int64, error)
	GetMoodSummary(string, time.Time, time.Time) (float64, int64, error)
}

type MoodRepository struct {
	db *gorm.DB
}

func NewMoodRepository(db *gorm.DB) IMoodRepository {
	return &MoodRepository{db}
}

func (r *MoodRepository) CreateMood(request *models.Mood) error {
	return r.db.CreateInBatches(request, 1000).Error
}

func (r *MoodRepository) GetUserMoods(userID string, params pagination.Params) ([]models.Mood, int64, error) {
	var moods []models.Mood
	var total int64

	if err := r.db.Model(&models.Mood{}).
		Where("id_user = ?", userID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Where("id_user = ?", userID).
		Order("date DESC").
		Limit(params.Limit).
		Offset(params.Skip).
		Find(&moods).Error

	return moods, total, err
}

func (r *MoodRepository) GetMoodSummary(userID string, startDate, endDate time.Time) (float64, int64, error) {
	var avg float64
	var count int64
	
	err := r.db.Model(&models.Mood{}).
		Where("id_user = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).
		Count(&count).Error
	if err != nil {
		return 0, 0, err
	}
	if count == 0 {
		return 0, 0, errors.New("data summary tidak ditemukan")
	}

	err = r.db.Model(&models.Mood{}).
		Select("AVG(mood_score)").
		Where("id_user = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).
		Scan(&avg).Error
	if err != nil {
		return 0, 0, err
	}


	return avg, count, nil
}