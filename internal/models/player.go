package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const (
	PositionForward    = "FORWARD"
	PositionMidfielder = "MIDFIELDER"
	PositionDefender   = "DEFENDER"
	PositionGoalkeeper = "GOALKEEPER"
)

type Player struct {
	ID           uint            `gorm:"primaryKey" json:"id"`
	TeamID       uint            `gorm:"not null;uniqueIndex:idx_team_jersey" json:"team_id"`
	Team         Team            `json:"team,omitempty"`
	Name         string          `gorm:"size:150;not null" json:"name"`
	HeightCM     decimal.Decimal `gorm:"type:decimal(5,2);not null" json:"height_cm"`
	WeightKG     decimal.Decimal `gorm:"type:decimal(5,2);not null" json:"weight_kg"`
	Position     string          `gorm:"size:20;not null" json:"position"`
	JerseyNumber int             `gorm:"not null;uniqueIndex:idx_team_jersey" json:"jersey_number"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `gorm:"index" json:"-"`
}
