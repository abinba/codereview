package model

import (
	"time"

	"github.com/google/uuid"
)

// CodeSnippetRating is a rating for a code snippet.
// @Description CodeSnippetRating is the model representing a rating for a code snippet.
type CodeSnippetRating struct {
	CodeSnippetRatingID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CodeSnippetVersionID uuid.UUID `gorm:"not null"`
	UserID              uuid.UUID `gorm:"not null"`
	Rating              int8	  `gorm:"not null"`
	CreatedAt		    time.Time `gorm:"autoCreateTime"`
	UpdatedAt		    time.Time `gorm:"autoUpdateTime"`
	User                User
}
