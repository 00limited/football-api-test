package repositories

import (
	"github.com/00limited/football-api/internal/models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) Create(admin *models.Admin) error {
	return r.db.Create(admin).Error
}

func (r *AdminRepository) GetByUsername(username string) (*models.Admin, error) {
	var admin models.Admin
	if err := r.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
