package models

import (
	"time"

	"gorm.io/gorm"
)

type MatchResult struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	MatchID   uint           `gorm:"not null;uniqueIndex" json:"match_id"`
	Match     Match          `json:"match,omitempty"`
	HomeScore int            `gorm:"not null" json:"home_score"`
	AwayScore int            `gorm:"not null" json:"away_score"`
	Goals     []Goal         `json:"goals,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
