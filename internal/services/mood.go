package services

import (
	"errors"
	"health-tech/internal/dto"
	"health-tech/internal/repository"
	"health-tech/models"
	"health-tech/pkg/pagination"
	"time"

	"github.com/google/uuid"
)

type IMoodService interface {
	CreateMood(string, *dto.CreateMood) error
	GetUserMoods(string, pagination.Params) ([]dto.DataMood, int, error)
	GetMoodSummary(string, string) (dto.GetSummary, error)
}

type MoodService struct {
	MoodRepository repository.IMoodRepository
	UserRepository repository.IUserRepository
}

func NewMoodService(moodRepository repository.IMoodRepository, userRepository repository.IUserRepository) IMoodService {
	return &MoodService{moodRepository, userRepository}
}

func (s *MoodService) CreateMood(userID string, request *dto.CreateMood) error {
	if request.Date.After(time.Now()) {
		return errors.New("tanggal tidak boleh di masa depan")
	}

	if request.MoodScore < 1 || request.MoodScore > 5 {
		return errors.New("mood score harus antara 1 dan 5")
	}

	if err := s.MoodRepository.CreateMood(&models.Mood{
		MoodID:    uuid.New().String(),
		IDUser:    userID,
		Date:      request.Date,
		MoodScore: request.MoodScore,
		MoodLabel: &request.MoodLabel,
		Notes:     &request.Notes,
	}); err != nil {
		return err
	}

	return nil
}

func (s *MoodService) GetUserMoods(userID string, params pagination.Params) ([]dto.DataMood, int, error) {
	existing, err := s.UserRepository.GetUser(dto.UserParams{UserID: userID})
	if err != nil {
		return nil, 0, err
	}
	if existing == nil {
		return nil, 0, errors.New("data tidak ditemukan")
	}
	
	moods, total, err := s.MoodRepository.GetUserMoods(userID, params)
	if err != nil {
		return nil, 0, err
	}

	var response []dto.DataMood
	for _, mood := range moods {
		response = append(response, dto.DataMood{
			MoodID:    mood.MoodID,
			Date:      mood.Date,
			MoodScore: mood.MoodScore,
			MoodLabel: *mood.MoodLabel,
			Notes:     *mood.Notes,
		})
	}

	return response, int(total), nil
}

func (s *MoodService) GetMoodSummary(userID string, period string) (dto.GetSummary, error) {
	now := time.Now()
	var startDate, endDate time.Time

	existing, err := s.UserRepository.GetUser(dto.UserParams{UserID: userID})
	if err != nil {
		return dto.GetSummary{}, err
	}
	if existing == nil {
		return dto.GetSummary{}, errors.New("data tidak ditemukan")
	}

	user, err := s.UserRepository.GetUser(dto.UserParams{UserID: userID})
	if err != nil {
		return dto.GetSummary{}, err
	}

	switch period {
	case "week":
		weekday := int(now.Weekday())
		startDate = now.AddDate(0, 0, -weekday)
		endDate = startDate.AddDate(0, 0, 7)
	case "month":
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endDate = startDate.AddDate(0, 1, 0)
	default:
		startDate = time.Time{}
		endDate = now
	}

	avg, count, err := s.MoodRepository.GetMoodSummary(userID, startDate, endDate)
	if err != nil {
		return dto.GetSummary{}, err
	}

	result := dto.GetSummary{
		User: dto.DataUser{
			UserID: userID,
			Nama:   user.Nama,
			Email:  user.Email,
		},
		Avg:       avg,
		TotalData: count,
	}
	return result, nil
}
