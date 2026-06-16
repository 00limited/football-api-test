package request

type CreatePlayerRequest struct {
	Name         string  `json:"name" validate:"required,min=2,max=150"`
	HeightCM     float64 `json:"height_cm" validate:"required,gt=0"`
	WeightKG     float64 `json:"weight_kg" validate:"required,gt=0"`
	Position     string  `json:"position" validate:"required,oneof=FORWARD MIDFIELDER DEFENDER GOALKEEPER"`
	JerseyNumber int     `json:"jersey_number" validate:"required,gte=1,lte=99"`
}

type UpdatePlayerRequest struct {
	Name         string  `json:"name" validate:"required,min=2,max=150"`
	HeightCM     float64 `json:"height_cm" validate:"required,gt=0"`
	WeightKG     float64 `json:"weight_kg" validate:"required,gt=0"`
	Position     string  `json:"position" validate:"required,oneof=FORWARD MIDFIELDER DEFENDER GOALKEEPER"`
	JerseyNumber int     `json:"jersey_number" validate:"required,gte=1,lte=99"`
}
