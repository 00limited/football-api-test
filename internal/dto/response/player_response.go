package response

type PlayerResponse struct {
	PlayerID  uint   `json:"player_id"`
	Name      string `json:"name"`
	Position  string `json:"position"`
	JerseyNum int    `json:"jersey_number"`
}

type PlayerDetailResponse struct {
	PlayerID  uint         `json:"player_id"`
	Name      string       `json:"name"`
	Position  string       `json:"position"`
	JerseyNum int          `json:"jersey_number"`
	Height    float64      `json:"height"`
	Weight    float64      `json:"weight"`
	Team      TeamResponse `json:"team"`
}
