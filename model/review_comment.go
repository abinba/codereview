package model

import (
	"time"

	"github.com/google/uuid"
)

// ReviewComment is a comment on a code snippet.
// @Description ReviewComment is the model representing a comment on a code snippet.
type ReviewComment struct {
	CommentID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID         *uuid.UUID
	CodeSnippetVersionID  uuid.UUID `gorm:"not null"`
	ReplyCommentID *uuid.UUID
	Text           string `gorm:"not null"`
	Line           *int
	IsGenerated    bool      `gorm:"default:false"`
	CreatedAt	   time.Time `gorm:"autoCreateTime"`
	UpdatedAt	   time.Time `gorm:"autoUpdateTime"`
	User           *User
	CodeSnippetVersion    CodeSnippetVersion
	ReplyComment   *ReviewComment
}
