package services

import (
	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/repositories"
)

type PlayerService struct {
	repo     *repositories.PlayerRepository
	teamRepo *repositories.TeamRepository
}

func NewPlayerService(repo *repositories.PlayerRepository, teamRepo *repositories.TeamRepository) *PlayerService {
	return &PlayerService{repo: repo, teamRepo: teamRepo}
}

func (s *PlayerService) Create(player *models.Player) error {
	if _, err := s.teamRepo.GetByID(player.TeamID); err != nil {
		return err
	}
	return s.repo.Create(player)
}

func (s *PlayerService) ListByTeamID(teamID uint) ([]models.Player, error) {
	if _, err := s.teamRepo.GetByID(teamID); err != nil {
		return nil, err
	}
	return s.repo.ListByTeamID(teamID)
}

func (s *PlayerService) GetByID(id uint) (*models.Player, error) { return s.repo.GetByID(id) }
func (s *PlayerService) Update(player *models.Player) error      { return s.repo.Update(player) }
func (s *PlayerService) Delete(player *models.Player) error      { return s.repo.Delete(player) }
