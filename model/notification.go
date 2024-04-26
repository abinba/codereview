package model

import (
	"time"

	"github.com/google/uuid"
)

// Notification is a model for representing notifications for particular user in the system.
// @Description Notification is a model for representing notifications for particular user in the system.
type Notification struct {
	NotificationID   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	NotificationType string    `gorm:"not null"`
	UserID           uuid.UUID `gorm:"not null"`
	Text             string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	User             User
}
