package server

import (
	"clypin/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	// Automigrate
	db.AutoMigrate(
		&models.User{},
		&models.Message{},
	)
}
