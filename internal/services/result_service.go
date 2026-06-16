package services

import (
	"fmt"

	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/repositories"
	"gorm.io/gorm"
)

type ResultService struct {
	matchRepo  *repositories.MatchRepository
	resultRepo *repositories.ResultRepository
	playerRepo *repositories.PlayerRepository
}

func NewResultService(matchRepo *repositories.MatchRepository, resultRepo *repositories.ResultRepository, playerRepo *repositories.PlayerRepository) *ResultService {
	return &ResultService{matchRepo: matchRepo, resultRepo: resultRepo, playerRepo: playerRepo}
}

func (s *ResultService) Create(matchID uint, goalRequests []models.Goal) (*models.MatchResult, error) {
	match, err := s.matchRepo.GetByID(matchID)
	if err != nil {
		return nil, err
	}
	if match.Status != models.MatchStatusScheduled {
		return nil, fmt.Errorf("match result can only be submitted for scheduled matches")
	}
	if exists, err := s.resultRepo.ExistsByMatchID(matchID); err != nil {
		return nil, err
	} else if exists {
		return nil, fmt.Errorf("match result already exists")
	}

	homeScore := 0
	awayScore := 0
	goals := make([]models.Goal, 0, len(goalRequests))
	for _, goal := range goalRequests {
		player, err := s.playerRepo.GetByID(goal.PlayerID)
		if err != nil {
			return nil, err
		}
		if player.TeamID != match.HomeTeamID && player.TeamID != match.AwayTeamID {
			return nil, fmt.Errorf("player %d does not belong to either team in the match", player.ID)
		}
		if player.TeamID == match.HomeTeamID {
			homeScore++
		} else {
			awayScore++
		}
		goals = append(goals, models.Goal{
			PlayerID:   goal.PlayerID,
			TeamID:     player.TeamID,
			GoalMinute: goal.GoalMinute,
		})
	}

	result := &models.MatchResult{MatchID: matchID, HomeScore: homeScore, AwayScore: awayScore}
	if err := s.resultRepo.Transaction(func(tx *gorm.DB) error {
		txResultRepo := repositories.NewResultRepository(tx)
		txMatchRepo := repositories.NewMatchRepository(tx)
		if err := txResultRepo.Create(result); err != nil {
			return err
		}
		for i := range goals {
			goals[i].MatchResultID = result.ID
		}
		if err := txResultRepo.CreateGoals(goals); err != nil {
			return err
		}
		match.Status = models.MatchStatusFinished
		if err := txMatchRepo.Update(match); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return s.resultRepo.GetByMatchID(matchID)
}

func (s *ResultService) GetByMatchID(matchID uint) (*models.MatchResult, error) {
	if _, err := s.matchRepo.GetByID(matchID); err != nil {
		return nil, err
	}
	return s.resultRepo.GetByMatchID(matchID)
}
