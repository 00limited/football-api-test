package repositories

import (
	"github.com/00limited/football-api/internal/models"
	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) ListFinishedMatches() ([]models.Match, error) {
	var matches []models.Match
	if err := r.db.
		Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("Result").
		Preload("Result.Goals").
		Preload("Result.Goals.Player").
		Preload("Result.Goals.Team").
		Where("status = ?", models.MatchStatusFinished).
		Order("match_date asc, match_time asc, id asc").
		Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

func (r *ReportRepository) GetFinishedMatchByID(id uint) (*models.Match, error) {
	var match models.Match
	if err := r.db.
		Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("Result").
		Preload("Result.Goals").
		Preload("Result.Goals.Player").
		Preload("Result.Goals.Team").
		Where("status = ?", models.MatchStatusFinished).
		First(&match, id).Error; err != nil {
		return nil, err
	}
	return &match, nil
}
