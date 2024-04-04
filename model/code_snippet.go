package model

import (
	"time"

	"github.com/google/uuid"
)

// CodeSnippet is a code snippet that is posted by a user.
// @Description CodeSnippet is the model representing a code snippet in the system.
type CodeSnippet struct {
	CodeSnippetID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID             uuid.UUID
	ProgramLanguageID  uuid.UUID
	Text               string
	IsPrivate          *bool `gorm:"default:false"`
	IsArchived         *bool `gorm:"default:false"`
	IsDraft            *bool `gorm:"default:false"`
	User               User
	CreatedAt		   time.Time `gorm:"autoCreateTime"`
	UpdatedAt		   time.Time `gorm:"autoUpdateTime"`
	ReviewComments     []ReviewComment
	CodeSnippetRatings []CodeSnippetRating
	ProgramLanguage    ProgramLanguage
}
