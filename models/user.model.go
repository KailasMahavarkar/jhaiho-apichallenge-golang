package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

)

type User struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}
