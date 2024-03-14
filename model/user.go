package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is a user in the system.
// @Description User is the model representing a user in the system.
type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;" description:"UUID"`
	Username string    `json:"username" example:"johndoe" description:"The username of the user"`
	Email    string    `json:"email" example:"johndoe@gmail.com" description:"The email of the user"`
	Password string    `json:"password" example:"12345678" description:"The password of the user"`
}

type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
