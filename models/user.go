package models

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"not null" json:"email"`
	FirstName string    `gorm:"not null" json:"first_name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	Password  string    `gorm:"not null" json:"password"`
	Messages  []Message `json:"messages"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) Create() User {
	user := User{
		Email:     "bustajay30@gmail.com",
		FirstName: "Bolu",
		LastName:  "Busta",
		Password:  "Mil",
	}
	// db.Create(&user)
	return user

}
