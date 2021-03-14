package domain

import (
	"time"

	"gorm.io/gorm"
)

// Base struct contain common fields for all domain struct
type Base struct {
	ID        ID             `json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"->"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"->"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
