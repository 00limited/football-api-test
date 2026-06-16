package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	MatchStatusScheduled = "SCHEDULED"
	MatchStatusFinished  = "FINISHED"
)

type Match struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	MatchDate  time.Time      `gorm:"type:date;not null" json:"match_date"`
	MatchTime  string         `gorm:"size:5;not null" json:"match_time"`
	HomeTeamID uint           `gorm:"not null" json:"home_team_id"`
	HomeTeam   Team           `gorm:"foreignKey:HomeTeamID" json:"home_team,omitempty"`
	AwayTeamID uint           `gorm:"not null" json:"away_team_id"`
	AwayTeam   Team           `gorm:"foreignKey:AwayTeamID" json:"away_team,omitempty"`
	Status     string         `gorm:"size:20;not null" json:"status"`
	Result     *MatchResult   `json:"result,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
