package request

type GoalRequest struct {
	PlayerID   uint `json:"player_id" validate:"required,gt=0"`
	GoalMinute int  `json:"goal_minute" validate:"required,gte=0,lte=130"`
}

type CreateResultRequest struct {
	Goals []GoalRequest `json:"goals" validate:"omitempty,dive"`
}
