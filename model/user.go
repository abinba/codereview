package model

import (
	"time"

	"github.com/google/uuid"
)

// User is a user in the system.
// @Description User is the model representing a user in the system.
type User struct {
	UserID   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" description:"UUID"`
	Username string    `json:"username" example:"johndoe" description:"The username of the user"`
	IsAnonymous bool   `gorm:"default:false" description:"Is the user anonymous"`
	IsActive    bool   `gorm:"default:true" description:"Is the user active"`
	CreatedAt		    time.Time `gorm:"autoCreateTime"`
	UpdatedAt		    time.Time `gorm:"autoUpdateTime"`
}

type Users struct {
	Users []User `json:"users"`
}
