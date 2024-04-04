package model

import (
	"github.com/google/uuid"
)

// ProgramLanguage is a programming language that a code snippet is written in.
// @Description ProgramLanguage is the model representing a programming language in the system.
type ProgramLanguage struct {
	ProgramLanguageID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name              string
}
