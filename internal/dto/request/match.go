package request

type CreateMatchRequest struct {
	MatchDate  string `json:"match_date" validate:"required,datetime=2006-01-02"`
	MatchTime  string `json:"match_time" validate:"required,datetime=15:04"`
	HomeTeamID uint   `json:"home_team_id" validate:"required,gt=0"`
	AwayTeamID uint   `json:"away_team_id" validate:"required,gt=0"`
}

type UpdateMatchRequest struct {
	MatchDate  string `json:"match_date" validate:"required,datetime=2006-01-02"`
	MatchTime  string `json:"match_time" validate:"required,datetime=15:04"`
	HomeTeamID uint   `json:"home_team_id" validate:"required,gt=0"`
	AwayTeamID uint   `json:"away_team_id" validate:"required,gt=0"`
}
