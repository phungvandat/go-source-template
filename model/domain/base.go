package domain

import (
	"time"
)

// Base struct contain common fields for all domain struct
type Base struct {
	ID ID `gorm:"type:uuid" json:"id"`
	// Created time
	CreatedAt time.Time `json:"created_at" gorm:"->"`
	// Updated time
	UpdatedAt *time.Time `json:"updated_at" gorm:"->"`
}
