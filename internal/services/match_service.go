package services

import (
	"fmt"
	"time"

	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/repositories"
)

type MatchService struct {
	repo     *repositories.MatchRepository
	teamRepo *repositories.TeamRepository
}

func NewMatchService(repo *repositories.MatchRepository, teamRepo *repositories.TeamRepository) *MatchService {
	return &MatchService{repo: repo, teamRepo: teamRepo}
}

func (s *MatchService) ParseDate(raw string) (time.Time, error) {
	return time.Parse("2006-01-02", raw)
}

func (s *MatchService) ValidateTeams(homeTeamID, awayTeamID uint) error {
	if homeTeamID == awayTeamID {
		return fmt.Errorf("home team and away team must be different")
	}
	if _, err := s.teamRepo.GetByID(homeTeamID); err != nil {
		return err
	}
	if _, err := s.teamRepo.GetByID(awayTeamID); err != nil {
		return err
	}
	return nil
}

func (s *MatchService) Create(match *models.Match) error {
	if err := s.ValidateTeams(match.HomeTeamID, match.AwayTeamID); err != nil {
		return err
	}
	match.Status = models.MatchStatusScheduled
	return s.repo.Create(match)
}

func (s *MatchService) List() ([]models.Match, error)          { return s.repo.List() }
func (s *MatchService) GetByID(id uint) (*models.Match, error) { return s.repo.GetByID(id) }

func (s *MatchService) Update(match *models.Match) error {
	if match.Status == models.MatchStatusFinished {
		return fmt.Errorf("finished matches cannot be updated")
	}
	if err := s.ValidateTeams(match.HomeTeamID, match.AwayTeamID); err != nil {
		return err
	}
	return s.repo.Update(match)
}

func (s *MatchService) Delete(match *models.Match) error {
	if match.Status == models.MatchStatusFinished {
		return fmt.Errorf("finished matches cannot be deleted")
	}
	return s.repo.Delete(match)
}
