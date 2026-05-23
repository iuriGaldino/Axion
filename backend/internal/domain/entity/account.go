package entity

import (
	"time"
)

type AccountType string

const (
	CheckingAccount AccountType = "checking"
	SavingsAccount  AccountType = "savings"
	CreditCard      AccountType = "credit_card"
	Investment      AccountType = "investment"
)

type Account struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	Name      string      `json:"name"`
	Type      AccountType `json:"type"`
	Balance   float64     `json:"balance"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
