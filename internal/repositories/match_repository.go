package repositories

import (
	"github.com/00limited/football-api/internal/models"
	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) Create(match *models.Match) error {
	return r.db.Create(match).Error
}

func (r *MatchRepository) List() ([]models.Match, error) {
	var matches []models.Match
	if err := r.db.Preload("HomeTeam").Preload("AwayTeam").Order("match_date asc, match_time asc, id asc").Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

func (r *MatchRepository) GetByID(id uint) (*models.Match, error) {
	var match models.Match
	if err := r.db.Preload("HomeTeam").Preload("AwayTeam").First(&match, id).Error; err != nil {
		return nil, err
	}
	return &match, nil
}

func (r *MatchRepository) Update(match *models.Match) error {
	return r.db.Save(match).Error
}

func (r *MatchRepository) Delete(match *models.Match) error {
	return r.db.Delete(match).Error
}
