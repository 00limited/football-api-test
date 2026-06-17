package response

type ScoreResponse struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type TopScorerResponse struct {
	PlayerID   uint   `json:"player_id"`
	PlayerName string `json:"player_name"`
	TeamID     uint   `json:"team_id"`
	TeamName   string `json:"team_name"`
	Goals      int    `json:"goals"`
}

type MatchReportResponse struct {
	MatchID                 uint               `json:"match_id"`
	MatchDate               string             `json:"match_date"`
	MatchTime               string             `json:"match_time"`
	HomeTeam                interface{}        `json:"home_team"`
	AwayTeam                interface{}        `json:"away_team"`
	FinalScore              ScoreResponse      `json:"final_score"`
	MatchStatus             string             `json:"match_status"`
	TopScorer               *TopScorerResponse `json:"top_scorer,omitempty"`
	AccumulatedHomeTeamWins int                `json:"accumulated_home_team_total_wins"`
	AccumulatedAwayTeamWins int                `json:"accumulated_away_team_total_wins"`
}
