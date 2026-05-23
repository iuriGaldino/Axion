package dto

import "time"

type CreateBudgetRequest struct {
	CategoryID string    `json:"category_id" validate:"required"`
	Amount     float64   `json:"amount" validate:"required"`
	Period     string    `json:"period"` // monthly, weekly
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date"`
}

type CreateGoalRequest struct {
	Name         string    `json:"name" validate:"required"`
	TargetAmount float64   `json:"target_amount" validate:"required"`
	Deadline     time.Time `json:"deadline"`
}

type GoalResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	TargetAmount  float64   `json:"target_amount"`
	CurrentAmount float64   `json:"current_amount"`
	Progress      float64   `json:"progress"`
	Deadline      time.Time `json:"deadline"`
}
