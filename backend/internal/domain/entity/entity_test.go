package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := NewUser("John Doe", "john@example.com", "hashed_password")
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "pt-BR", user.Preferences.Language)
	assert.Equal(t, "system", user.Preferences.Theme)
	assert.NotEmpty(t, user.ID)
}

func TestGoal_CalculateProgress(t *testing.T) {
	goal := NewGoal(uuid.New(), "Save for car", 10000, time.Now().AddDate(0, 6, 0))

	assert.Equal(t, 0.0, goal.CalculateProgress())

	goal.CurrentAmount = 5000
	assert.Equal(t, 50.0, goal.CalculateProgress())

	goal.CurrentAmount = 10000
	assert.Equal(t, 100.0, goal.CalculateProgress())

	goal.CurrentAmount = 15000
	assert.Equal(t, 100.0, goal.CalculateProgress())
}
