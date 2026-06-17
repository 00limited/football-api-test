package response

type TeamResponse struct {
	TeamID      uint   `json:"team_id"`
	Name        string `json:"name"`
	LogoURL     string `json:"logo_url"`
	FoundedYear int    `json:"founded_year"`
	Address     string `json:"address"`
	City        string `json:"city"`
}

type TeamDetailResponse struct {
	TeamID      uint             `json:"team_id"`
	Name        string           `json:"name"`
	LogoURL     string           `json:"logo_url"`
	FoundedYear int              `json:"founded_year"`
	Address     string           `json:"address"`
	City        string           `json:"city"`
	Players     []PlayerResponse `json:"players"`
}
