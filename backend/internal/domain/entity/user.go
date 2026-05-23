package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserPreferences struct {
	Language    string `json:"language"`
	Theme       string `json:"theme"`
	AccentColor string `json:"accent_color"`
}

type User struct {
	ID           uuid.UUID       `json:"id"`
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	PasswordHash string          `json:"-"`
	Avatar       string          `json:"avatar"`
	Preferences  UserPreferences `json:"preferences"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

func NewUser(name, email, passwordHash string) *User {
	now := time.Now()
	return &User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Preferences: UserPreferences{
			Language: "pt-BR",
			Theme:    "system",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
}
