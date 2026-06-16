package models

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:150;not null;uniqueIndex" json:"name"`
	LogoURL     string         `gorm:"size:255" json:"logo_url"`
	FoundedYear int            `gorm:"not null" json:"founded_year"`
	Address     string         `gorm:"size:255;not null" json:"address"`
	City        string         `gorm:"size:100;not null" json:"city"`
	Players     []Player       `json:"players,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
