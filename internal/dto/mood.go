package dto

import (
	"health-tech/pkg/pagination"
	"time"
)

type CreateMood struct {
	UserID    string    `json:"user_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	MoodScore int       `json:"mood_score" binding:"required,min=1,max=5"`
	MoodLabel string    `json:"mood_label"`
	Notes     string    `json:"notes"`
}

type DataMood struct {
	MoodID    string    `json:"mood_id"`
	Date      time.Time `json:"date"`
	MoodScore int       `json:"mood_score"`
	MoodLabel string    `json:"mood_label"`
	Notes     string    `json:"notes"`
}

type DataMoodPagination struct {
	Meta  pagination.Meta `json:"meta"`
	Moods []DataMood      `json:"moods"`
}

type GetMoodUserPagination struct {
	User  DataUser             `json:"data_user"`
	Moods []DataMoodPagination `json:"moods"`
}

type GetSummary struct {
	User      DataUser `json:"data_user"`
	Avg       float64  `json:"rata_rata"`
	TotalData int64    `json:"totol_data"`
}
