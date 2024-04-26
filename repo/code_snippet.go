package repo

import (
	"github.com/abinba/codereview/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CodeSnippetRepository struct {
	db *gorm.DB
}

func NewCodeSnippetRepository(db *gorm.DB) *CodeSnippetRepository {
	return &CodeSnippetRepository{db: db}
}

func (repo *CodeSnippetRepository) CreateCodeSnippet(user_id *uuid.UUID, title string, is_private *bool) error {
	user := model.CodeSnippet{
		UserID:    user_id,
		Title:     title,
		IsPrivate: is_private,
	}
	return repo.db.Create(&user).Error
}
