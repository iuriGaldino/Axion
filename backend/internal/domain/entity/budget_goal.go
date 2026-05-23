package entity

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	CategoryID uuid.UUID `json:"category_id"`
	Amount     float64   `json:"amount"`
	Period     string    `json:"period"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewBudget(userID, categoryID uuid.UUID, amount float64, period string, start, end time.Time) *Budget {
	return &Budget{
		ID:         uuid.New(),
		UserID:     userID,
		CategoryID: categoryID,
		Amount:     amount,
		Period:     period,
		StartDate:  start,
		EndDate:    end,
		CreatedAt:  time.Now(),
	}
}

type Goal struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Name          string    `json:"name"`
	TargetAmount  float64   `json:"target_amount"`
	CurrentAmount float64   `json:"current_amount"`
	Deadline      time.Time `json:"deadline"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NewGoal(userID uuid.UUID, name string, target float64, deadline time.Time) *Goal {
	now := time.Now()
	return &Goal{
		ID:            uuid.New(),
		UserID:        userID,
		Name:          name,
		TargetAmount:  target,
		CurrentAmount: 0,
		Deadline:      deadline,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

func (g *Goal) CalculateProgress() float64 {
	if g.TargetAmount <= 0 {
		return 0
	}
	progress := (g.CurrentAmount / g.TargetAmount) * 100
	if progress > 100 {
		return 100
	}
	return progress
}
