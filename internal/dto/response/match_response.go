package response

type MatchResponse struct {
	MatchID   uint         `json:"match_id"`
	MatchDate string       `json:"match_date"`
	MatchTime string       `json:"match_time"`
	Status    string       `json:"status"`
	Result    string       `json:"result,omitempty"`
	HomeTeam  TeamResponse `json:"home_team"`
	AwayTeam  TeamResponse `json:"away_team"`
}

type MatchDetailResponse struct {
	MatchID   uint                 `json:"match_id"`
	MatchDate string               `json:"match_date"`
	MatchTime string               `json:"match_time"`
	Status    string               `json:"status"`
	HomeTeam  TeamDetailResponse   `json:"home_team"`
	AwayTeam  TeamDetailResponse   `json:"away_team"`
	Result    *MatchResultResponse `json:"result,omitempty"`
}

type MatchResultResponse struct {
	HomeScore int `json:"home_score"`
	AwayScore int `json:"away_score"`
}
