package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:100;uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"not null" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
