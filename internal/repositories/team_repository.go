package repositories

import (
	"github.com/00limited/football-api/internal/models"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) Create(team *models.Team) error {
	return r.db.Create(team).Error
}

func (r *TeamRepository) List() ([]models.Team, error) {
	var teams []models.Team
	if err := r.db.Order("name asc").Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *TeamRepository) GetByID(id uint) (*models.Team, error) {
	var team models.Team
	if err := r.db.First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *TeamRepository) Update(team *models.Team) error {
	return r.db.Save(team).Error
}

func (r *TeamRepository) Delete(team *models.Team) error {
	return r.db.Delete(team).Error
}
