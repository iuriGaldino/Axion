package entity

import (
	"time"
	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"` // checking, savings, credit_card, cash
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"` // ISO 4217
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccount(userID uuid.UUID, name, aType, currency string) *Account {
	now := time.Now()
	return &Account{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		Type:      aType,
		Balance:   0,
		Currency:  currency,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
