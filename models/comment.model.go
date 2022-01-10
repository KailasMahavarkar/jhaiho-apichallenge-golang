package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Text     string    `json:"text"`
}
