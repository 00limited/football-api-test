package repositories

import (
	"github.com/00limited/football-api/internal/models"
	"gorm.io/gorm"
)

type ResultRepository struct {
	db *gorm.DB
}

func NewResultRepository(db *gorm.DB) *ResultRepository {
	return &ResultRepository{db: db}
}

func (r *ResultRepository) Create(result *models.MatchResult) error {
	return r.db.Create(result).Error
}

func (r *ResultRepository) CreateGoals(goals []models.Goal) error {
	if len(goals) == 0 {
		return nil
	}
	return r.db.Create(&goals).Error
}

func (r *ResultRepository) GetByMatchID(matchID uint) (*models.MatchResult, error) {
	var result models.MatchResult
	if err := r.db.
		Preload("Goals", func(db *gorm.DB) *gorm.DB { return db.Order("goal_minute asc, id asc") }).
		Preload("Goals.Player").
		Preload("Goals.Team").
		Preload("Match.HomeTeam").
		Preload("Match.AwayTeam").
		First(&result, "match_id = ?", matchID).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ResultRepository) ExistsByMatchID(matchID uint) (bool, error) {
	var count int64
	if err := r.db.Model(&models.MatchResult{}).Where("match_id = ?", matchID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *ResultRepository) Transaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}
