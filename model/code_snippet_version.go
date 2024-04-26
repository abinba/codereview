package model

import (
	"time"

	"github.com/google/uuid"
)

// CodeSnippetVersion is a code snippet version that is posted by a user.
// @Description CodeSnippetVersion is the model representing a code snippet version of the code snippet in the system.
type CodeSnippetVersion struct {
	CodeSnippetVersionID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CodeSnippetID        uuid.UUID `gorm:"not null"`
	ProgramLanguageID    uuid.UUID `gorm:"not null"`
	Text                 string    `gorm:"not null"`
	CreatedAt            time.Time `gorm:"autoCreateTime"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime"`
	ProgramLanguage      ProgramLanguage
	CodeSnippetRatings   []CodeSnippetRating
	ReviewComments       []ReviewComment
}
