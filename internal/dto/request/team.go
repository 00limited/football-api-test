package request

type CreateTeamRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=150"`
	FoundedYear int    `json:"founded_year" validate:"required,gte=1800,lte=3000"`
	Address     string `json:"address" validate:"required,max=255"`
	City        string `json:"city" validate:"required,max=100"`
}

type UpdateTeamRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=150"`
	FoundedYear int    `json:"founded_year" validate:"required,gte=1800,lte=3000"`
	Address     string `json:"address" validate:"required,max=255"`
	City        string `json:"city" validate:"required,max=100"`
}
