package model

import (
	"time"

	"github.com/google/uuid"
)

// CodeSnippetRating is a rating for a code snippet.
// @Description CodeSnippetRating is the model representing a rating for a code snippet.
type CodeSnippetRating struct {
	CodeSnippetRatingID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CodeSnippetID       uuid.UUID
	UserID              uuid.UUID
	Rating              int8
	CreatedAt		    time.Time `gorm:"autoCreateTime"`
	UpdatedAt		    time.Time `gorm:"autoUpdateTime"`
	User                User
	CodeSnippet         CodeSnippet
}
