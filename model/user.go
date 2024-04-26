package model

import (
	"time"

	"github.com/google/uuid"
)

// User is a user in the system.
// @Description User is the model representing a user in the system.
type User struct {
	UserID   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" description:"UUID"`
	Username string    `gorm:"unqiue;not null" json:"username" example:"johndoe" description:"The username of the user"`
	Password string	   `gorm:"not null;" json:"password" description:"The password of the user"`
	IsActive    bool   `gorm:"default:true" description:"Is the user active"`
	CreatedAt		    time.Time `gorm:"autoCreateTime"`
	UpdatedAt		    time.Time `gorm:"autoUpdateTime"`
}

type Users struct {
	Users []User `json:"users"`
}
