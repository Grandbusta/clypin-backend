package queries

import (
	"clypin/config"
	"clypin/models"
)

func CreateUser(user *models.User) (*models.User, error) {
	db := config.DB
	err := db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func FindUserByEmail(email string) (*models.User, error) {
	db := config.DB
	user := models.User{}
	err := db.Model(&models.User{}).First(&user, models.User{Email: email}).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}
