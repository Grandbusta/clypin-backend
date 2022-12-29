package models

import "time"

type User struct {
	ID        uint64
	Email     string
	FirstName string
	LastName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Create() {

}
