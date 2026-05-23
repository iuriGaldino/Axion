package entity

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	IsIncome  bool      `json:"is_income"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCategory(userID uuid.UUID, name, icon, color string, isIncome bool) *Category {
	return &Category{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		Icon:      icon,
		Color:     color,
		IsIncome:  isIncome,
		CreatedAt: time.Now(),
	}
}
