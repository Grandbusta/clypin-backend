package models

import "time"

type Message struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	User_ID   uint64    `json:"user_Id"`
	Content   string    `gorm:"not null" json:"content"`
	Client    string    `gorm:"not null" json:"client"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
