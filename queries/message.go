package queries

import (
	"clypin/config"
	"clypin/models"
)

func CreateMessage(msg *models.Message) *models.Message {
	db := config.DB
	db.Create(&msg)
	return msg
}

func FetchMessages(userID uint64) (*[]models.Message, error) {
	db := config.DB
	msgs := []models.Message{}
	err := db.Model(&models.Message{}).Find(&msgs, models.Message{UserID: userID}).Error
	if err != nil {
		return &msgs, err
	}
	return &msgs, nil
}
