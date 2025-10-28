package models

import "time"

type Mood struct {
	MoodID    string    `json:"mood_id" gorm:"type:char(36);primaryKey"`
	IDUser    string    `json:"user_id" gorm:"type:char(36);not null;index"`
	Date      time.Time `json:"date" gorm:"not null;index"`
	MoodScore int       `json:"mood_score" gorm:"not null;check:mood_score >= 1 AND mood_score <= 5"`
	MoodLabel *string   `json:"mood_label" gorm:"type:varchar(100)"`
	Notes     *string   `json:"notes" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
