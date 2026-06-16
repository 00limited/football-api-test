package repositories

import (
	"github.com/00limited/football-api/internal/models"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Create(player *models.Player) error {
	return r.db.Create(player).Error
}

func (r *PlayerRepository) ListByTeamID(teamID uint) ([]models.Player, error) {
	var players []models.Player
	if err := r.db.Where("team_id = ?", teamID).Order("jersey_number asc").Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (r *PlayerRepository) GetByID(id uint) (*models.Player, error) {
	var player models.Player
	if err := r.db.First(&player, id).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) Update(player *models.Player) error {
	return r.db.Save(player).Error
}

func (r *PlayerRepository) Delete(player *models.Player) error {
	return r.db.Delete(player).Error
}
