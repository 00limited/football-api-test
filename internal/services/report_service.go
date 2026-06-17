package services

import (
	"sort"

	resp "github.com/00limited/football-api/internal/dto/response"
	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) List() ([]resp.MatchReportResponse, error) {
	matches, err := s.repo.ListFinishedMatches()
	if err != nil {
		return nil, err
	}
	return s.buildReports(matches), nil
}

func (s *ReportService) GetByMatchID(id uint) (*resp.MatchReportResponse, error) {
	matches, err := s.repo.ListFinishedMatches()
	if err != nil {
		return nil, err
	}
	reports := s.buildReports(matches)
	for _, report := range reports {
		if report.MatchID == id {
			return &report, nil
		}
	}
	return nil, gormNotFound()
}

func (s *ReportService) buildReports(matches []models.Match) []resp.MatchReportResponse {
	wins := map[uint]int{}
	reports := make([]resp.MatchReportResponse, 0, len(matches))
	for _, match := range matches {
		status := reportStatus(match)
		switch status {
		case "HOME_WIN":
			wins[match.HomeTeamID]++
		case "AWAY_WIN":
			wins[match.AwayTeamID]++
		}
		reports = append(reports, resp.MatchReportResponse{
			MatchID:   match.ID,
			MatchDate: match.MatchDate.Format("2006-01-02"),
			MatchTime: match.MatchTime,
			HomeTeam: map[string]interface{}{
				"id":   match.HomeTeam.ID,
				"name": match.HomeTeam.Name,
				"city": match.HomeTeam.City,
			},
			AwayTeam: map[string]interface{}{
				"id":   match.AwayTeam.ID,
				"name": match.AwayTeam.Name,
				"city": match.AwayTeam.City,
			},
			FinalScore: resp.ScoreResponse{
				Home: match.Result.HomeScore,
				Away: match.Result.AwayScore,
			},
			MatchStatus:             status,
			TopScorer:               topScorer(match),
			AccumulatedHomeTeamWins: wins[match.HomeTeamID],
			AccumulatedAwayTeamWins: wins[match.AwayTeamID],
		})
	}
	return reports
}

func reportStatus(match models.Match) string {
	if match.Result.HomeScore > match.Result.AwayScore {
		return "HOME_WIN"
	}
	if match.Result.AwayScore > match.Result.HomeScore {
		return "AWAY_WIN"
	}
	return "DRAW"
}

func topScorer(match models.Match) *resp.TopScorerResponse {
	if match.Result == nil || len(match.Result.Goals) == 0 {
		return nil
	}
	counts := map[uint]*resp.TopScorerResponse{}
	for _, goal := range match.Result.Goals {
		item, ok := counts[goal.PlayerID]
		if !ok {
			item = &resp.TopScorerResponse{
				PlayerID:   goal.PlayerID,
				PlayerName: goal.Player.Name,
				TeamID:     goal.TeamID,
				TeamName:   goal.Team.Name,
			}
			counts[goal.PlayerID] = item
		}
		item.Goals++
	}
	players := make([]*resp.TopScorerResponse, 0, len(counts))
	for _, item := range counts {
		players = append(players, item)
	}
	sort.Slice(players, func(i, j int) bool {
		if players[i].Goals == players[j].Goals {
			return players[i].PlayerID < players[j].PlayerID
		}
		return players[i].Goals > players[j].Goals
	})
	return players[0]
}
