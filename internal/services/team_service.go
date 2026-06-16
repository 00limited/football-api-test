package services

import (
	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/repositories"
)

type TeamService struct {
	repo *repositories.TeamRepository
}

func NewTeamService(repo *repositories.TeamRepository) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) Create(team *models.Team) error        { return s.repo.Create(team) }
func (s *TeamService) List() ([]models.Team, error)          { return s.repo.List() }
func (s *TeamService) GetByID(id uint) (*models.Team, error) { return s.repo.GetByID(id) }
func (s *TeamService) Update(team *models.Team) error        { return s.repo.Update(team) }
func (s *TeamService) Delete(team *models.Team) error        { return s.repo.Delete(team) }
