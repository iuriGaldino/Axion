package postgres

import (
	"context"
	"database/sql"

	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

type PostgresTransactionRepository struct {
	db *sql.DB
}

func NewPostgresTransactionRepository(db *sql.DB) *PostgresTransactionRepository {
	return &PostgresTransactionRepository{db: db}
}

func (r *PostgresTransactionRepository) Create(ctx context.Context, t *entity.Transaction) error {
	query := `
		INSERT INTO transactions (id, user_id, account_id, category_id, amount, description, date, type, is_confirmed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.ExecContext(ctx, query,
		t.ID, t.UserID, t.AccountID, t.CategoryID, t.Amount, t.Description, t.Date, t.Type, t.IsConfirmed, t.CreatedAt, t.UpdatedAt,
	)
	return err
}

func (r *PostgresTransactionRepository) GetByID(ctx context.Context, id string) (*entity.Transaction, error) {
	query := `SELECT id, user_id, account_id, category_id, amount, description, date, type, is_confirmed, created_at, updated_at FROM transactions WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var t entity.Transaction
	if err := row.Scan(&t.ID, &t.UserID, &t.AccountID, &t.CategoryID, &t.Amount, &t.Description, &t.Date, &t.Type, &t.IsConfirmed, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *PostgresTransactionRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Transaction, error) {
	query := `SELECT id, user_id, account_id, category_id, amount, description, date, type, is_confirmed, created_at, updated_at FROM transactions WHERE user_id = $1 ORDER BY date DESC`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*entity.Transaction
	for rows.Next() {
		var t entity.Transaction
		if err := rows.Scan(&t.ID, &t.UserID, &t.AccountID, &t.CategoryID, &t.Amount, &t.Description, &t.Date, &t.Type, &t.IsConfirmed, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, &t)
	}
	return transactions, nil
}

func (r *PostgresTransactionRepository) Update(ctx context.Context, t *entity.Transaction) error {
	query := `
		UPDATE transactions
		SET account_id = $1, category_id = $2, amount = $3, description = $4, date = $5, type = $6, is_confirmed = $7, updated_at = $8
		WHERE id = $9
	`
	_, err := r.db.ExecContext(ctx, query,
		t.AccountID, t.CategoryID, t.Amount, t.Description, t.Date, t.Type, t.IsConfirmed, t.UpdatedAt, t.ID,
	)
	return err
}

func (r *PostgresTransactionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM transactions WHERE id = $1", id)
	return err
}
