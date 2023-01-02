package queries

import (
	"clypin/config"
	"clypin/models"
)

func Create(user *models.User) *models.User {
	db := config.DB
	db.Create(&user)
	return user
}
