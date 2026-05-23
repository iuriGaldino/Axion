package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users (id, name, email, password_hash, created_at, updated_at, language, theme)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
		user.Preferences.Language,
		user.Preferences.Theme,
	)
	return err
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, name, email, password_hash, avatar, language, theme, accent_color, created_at, updated_at
		FROM users WHERE email = $1
	`
	row := r.db.QueryRowContext(ctx, query, email)

	var user entity.User
	var avatar, accentColor sql.NullString
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&avatar,
		&user.Preferences.Language,
		&user.Preferences.Theme,
		&accentColor,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	user.Avatar = avatar.String
	user.Preferences.AccentColor = accentColor.String

	return &user, nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	query := `
		SELECT id, name, email, password_hash, avatar, language, theme, accent_color, created_at, updated_at
		FROM users WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var user entity.User
	var avatar, accentColor sql.NullString
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&avatar,
		&user.Preferences.Language,
		&user.Preferences.Theme,
		&accentColor,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	user.Avatar = avatar.String
	user.Preferences.AccentColor = accentColor.String

	return &user, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *entity.User) error {
	query := `
		UPDATE users
		SET name = $1, avatar = $2, language = $3, theme = $4, accent_color = $5, updated_at = $6
		WHERE id = $7
	`
	_, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.Avatar,
		user.Preferences.Language,
		user.Preferences.Theme,
		user.Preferences.AccentColor,
		user.UpdatedAt,
		user.ID,
	)
	return err
}
