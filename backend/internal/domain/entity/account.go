package entity

import (
	"time"

	"github.com/google/uuid"
)

type AccountType string

const (
	CheckingAccount AccountType = "checking"
	SavingsAccount  AccountType = "savings"
	CreditCard      AccountType = "credit_card"
	Investment      AccountType = "investment"
)

type Account struct {
	ID        uuid.UUID   `json:"id"`
	UserID    uuid.UUID   `json:"user_id"`
	Name      string      `json:"name"`
	Type      AccountType `json:"type"`
	Balance   float64     `json:"balance"`
	Currency  string      `json:"currency"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func NewAccount(
	userID uuid.UUID,
	name string,
	accountType AccountType,
	currency string,
) *Account {
	now := time.Now()

	return &Account{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		Type:      accountType,
		Balance:   0,
		Currency:  currency,
		CreatedAt: now,
		UpdatedAt: now,
	}
}