package dto

import "time"

type CreateTransactionRequest struct {
	AccountID   string    `json:"account_id" validate:"required"`
	CategoryID  string    `json:"category_id"`
	Amount      float64   `json:"amount" validate:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" validate:"required"`
	Type        string    `json:"type" validate:"required"` // income, expense, transfer
}

type TransactionResponse struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	CategoryID  string    `json:"category_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Type        string    `json:"type"`
	IsConfirmed bool      `json:"is_confirmed"`
}
