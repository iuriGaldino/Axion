package postgres

import (
	"context"
	"database/sql"

	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

type PostgresBudgetRepository struct {
	db *sql.DB
}

func NewPostgresBudgetRepository(db *sql.DB) *PostgresBudgetRepository {
	return &PostgresBudgetRepository{db: db}
}

func (r *PostgresBudgetRepository) Create(ctx context.Context, b *entity.Budget) error {
	query := `INSERT INTO budgets (id, user_id, category_id, amount, period, start_date, end_date, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, b.ID, b.UserID, b.CategoryID, b.Amount, b.Period, b.StartDate, b.EndDate, b.CreatedAt)
	return err
}

func (r *PostgresBudgetRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Budget, error) {
	query := `SELECT id, user_id, category_id, amount, period, start_date, end_date, created_at FROM budgets WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var budgets []*entity.Budget
	for rows.Next() {
		var b entity.Budget
		if err := rows.Scan(&b.ID, &b.UserID, &b.CategoryID, &b.Amount, &b.Period, &b.StartDate, &b.EndDate, &b.CreatedAt); err != nil {
			return nil, err
		}
		budgets = append(budgets, &b)
	}
	return budgets, nil
}

func (r *PostgresBudgetRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM budgets WHERE id = $1", id)
	return err
}

type PostgresGoalRepository struct {
	db *sql.DB
}

func NewPostgresGoalRepository(db *sql.DB) *PostgresGoalRepository {
	return &PostgresGoalRepository{db: db}
}

func (r *PostgresGoalRepository) Create(ctx context.Context, g *entity.Goal) error {
	query := `INSERT INTO goals (id, user_id, name, target_amount, current_amount, deadline, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, g.ID, g.UserID, g.Name, g.TargetAmount, g.CurrentAmount, g.Deadline, g.CreatedAt, g.UpdatedAt)
	return err
}

func (r *PostgresGoalRepository) GetByID(ctx context.Context, id string) (*entity.Goal, error) {
	query := `SELECT id, user_id, name, target_amount, current_amount, deadline, created_at, updated_at FROM goals WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var g entity.Goal
	if err := row.Scan(&g.ID, &g.UserID, &g.Name, &g.TargetAmount, &g.CurrentAmount, &g.Deadline, &g.CreatedAt, &g.UpdatedAt); err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *PostgresGoalRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Goal, error) {
	query := `SELECT id, user_id, name, target_amount, current_amount, deadline, created_at, updated_at FROM goals WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []*entity.Goal
	for rows.Next() {
		var g entity.Goal
		if err := rows.Scan(&g.ID, &g.UserID, &g.Name, &g.TargetAmount, &g.CurrentAmount, &g.Deadline, &g.CreatedAt, &g.UpdatedAt); err != nil {
			return nil, err
		}
		goals = append(goals, &g)
	}
	return goals, nil
}

func (r *PostgresGoalRepository) Update(ctx context.Context, g *entity.Goal) error {
	query := `UPDATE goals SET name = $1, target_amount = $2, current_amount = $3, deadline = $4, updated_at = $5 WHERE id = $6`
	_, err := r.db.ExecContext(ctx, query, g.Name, g.TargetAmount, g.CurrentAmount, g.Deadline, g.UpdatedAt, g.ID)
	return err
}

func (r *PostgresGoalRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM goals WHERE id = $1", id)
	return err
}
