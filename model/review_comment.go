package model

import (
	"time"

	"github.com/google/uuid"
)

// ReviewComment is a comment on a code snippet.
// @Description ReviewComment is the model representing a comment on a code snippet.
type ReviewComment struct {
	CommentID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID         uuid.UUID
	CodeSnippetID  uuid.UUID
	ReplyCommentID uuid.UUID
	Text           string
	Line           int
	CreatedAt	   time.Time `gorm:"autoCreateTime"`
	UpdatedAt	   time.Time `gorm:"autoUpdateTime"`
	User           User
	CodeSnippet    CodeSnippet
}
