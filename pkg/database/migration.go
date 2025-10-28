package database

import (
	"health-tech/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Mood{},
	)
	if err != nil {
		return err
	}

	return nil
}