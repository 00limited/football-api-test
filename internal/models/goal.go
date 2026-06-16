package models

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	MatchResultID uint           `gorm:"not null" json:"match_result_id"`
	PlayerID      uint           `gorm:"not null" json:"player_id"`
	Player        Player         `json:"player,omitempty"`
	TeamID        uint           `gorm:"not null" json:"team_id"`
	Team          Team           `json:"team,omitempty"`
	GoalMinute    int            `gorm:"not null" json:"goal_minute"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
