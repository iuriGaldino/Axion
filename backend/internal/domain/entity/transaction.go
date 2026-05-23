package entity

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string

const (
	Income   TransactionType = "income"
	Expense  TransactionType = "expense"
	Transfer TransactionType = "transfer"
)

type Transaction struct {
	ID          uuid.UUID       `json:"id"`
	UserID      uuid.UUID       `json:"user_id"`
	AccountID   uuid.UUID       `json:"account_id"`
	CategoryID  *uuid.UUID      `json:"category_id"`
	Amount      float64         `json:"amount"`
	Description string          `json:"description"`
	Date        time.Time       `json:"date"`
	Type        TransactionType `json:"type"`
	IsConfirmed bool            `json:"is_confirmed"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

func NewTransaction(userID, accountID uuid.UUID, categoryID *uuid.UUID, amount float64, description string, date time.Time, tType TransactionType) *Transaction {
	now := time.Now()
	return &Transaction{
		ID:          uuid.New(),
		UserID:      userID,
		AccountID:   accountID,
		CategoryID:  categoryID,
		Amount:      amount,
		Description: description,
		Date:        date,
		Type:        tType,
		IsConfirmed: true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
