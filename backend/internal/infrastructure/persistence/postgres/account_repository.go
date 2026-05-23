package postgres

import (
	"context"
	"database/sql"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

type PostgresAccountRepository struct {
	db *sql.DB
}

func NewPostgresAccountRepository(db *sql.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{db: db}
}

func (r *PostgresAccountRepository) Create(ctx context.Context, a *entity.Account) error {
	query := `INSERT INTO accounts (id, user_id, name, type, balance, currency, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, a.ID, a.UserID, a.Name, a.Type, a.Balance, a.Currency, a.CreatedAt, a.UpdatedAt)
	return err
}

func (r *PostgresAccountRepository) GetByID(ctx context.Context, id string) (*entity.Account, error) {
	query := `SELECT id, user_id, name, type, balance, currency, created_at, updated_at FROM accounts WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	var a entity.Account
	if err := row.Scan(&a.ID, &a.UserID, &a.Name, &a.Type, &a.Balance, &a.Currency, &a.CreatedAt, &a.UpdatedAt); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *PostgresAccountRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Account, error) {
	query := `SELECT id, user_id, name, type, balance, currency, created_at, updated_at FROM accounts WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []*entity.Account
	for rows.Next() {
		var a entity.Account
		if err := rows.Scan(&a.ID, &a.UserID, &a.Name, &a.Type, &a.Balance, &a.Currency, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, &a)
	}
	return accounts, nil
}

func (r *PostgresAccountRepository) Update(ctx context.Context, a *entity.Account) error {
	query := `UPDATE accounts SET name = $1, type = $2, balance = $3, updated_at = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, a.Name, a.Type, a.Balance, a.UpdatedAt, a.ID)
	return err
}

func (r *PostgresAccountRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM accounts WHERE id = $1", id)
	return err
}
