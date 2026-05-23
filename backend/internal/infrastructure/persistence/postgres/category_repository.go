package postgres

import (
	"context"
	"database/sql"

	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

type PostgresCategoryRepository struct {
	db *sql.DB
}

func NewPostgresCategoryRepository(db *sql.DB) *PostgresCategoryRepository {
	return &PostgresCategoryRepository{db: db}
}

func (r *PostgresCategoryRepository) Create(ctx context.Context, c *entity.Category) error {
	query := `INSERT INTO categories (id, user_id, name, icon, color, is_income, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, c.ID, c.UserID, c.Name, c.Icon, c.Color, c.IsIncome, c.CreatedAt)
	return err
}

func (r *PostgresCategoryRepository) GetByID(ctx context.Context, id string) (*entity.Category, error) {
	query := `SELECT id, user_id, name, icon, color, is_income, created_at FROM categories WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var c entity.Category
	var icon, color sql.NullString
	if err := row.Scan(&c.ID, &c.UserID, &c.Name, &icon, &color, &c.IsIncome, &c.CreatedAt); err != nil {
		return nil, err
	}
	c.Icon = icon.String
	c.Color = color.String
	return &c, nil
}

func (r *PostgresCategoryRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Category, error) {
	query := `SELECT id, user_id, name, icon, color, is_income, created_at FROM categories WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var c entity.Category
		var icon, color sql.NullString
		if err := rows.Scan(&c.ID, &c.UserID, &c.Name, &icon, &color, &c.IsIncome, &c.CreatedAt); err != nil {
			return nil, err
		}
		c.Icon = icon.String
		c.Color = color.String
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *PostgresCategoryRepository) Update(ctx context.Context, c *entity.Category) error {
	query := `UPDATE categories SET name = $1, icon = $2, color = $3, is_income = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, c.Name, c.Icon, c.Color, c.IsIncome, c.ID)
	return err
}

func (r *PostgresCategoryRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM categories WHERE id = $1", id)
	return err
}
