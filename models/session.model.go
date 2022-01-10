package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Username string    `json:"username"`
}
