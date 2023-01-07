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

func FindByEmail(email string) (*models.User, error) {
	db := config.DB
	user := models.User{}
	err := db.Model(&models.User{}).First(&user, models.User{Email: email}).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}
