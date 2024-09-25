package application_test

import (
	"context"
	"testing"

	"github.com/hossein-225/Library-Management/auth-service/internal/application"
	"github.com/hossein-225/Library-Management/auth-service/internal/domain"
	"github.com/hossein-225/Library-Management/auth-service/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	service := application.NewAuthService(nil)

	token, err := service.GenerateToken(context.Background(), "user123")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := utils.ParseJWT(token)
	assert.NoError(t, err)
	assert.Equal(t, "user123", claims["userID"])
}

func TestValidateToken(t *testing.T) {
	repo := &MockAuthRepository{
		ValidToken: "valid_token",
		UserID:     "user123",
	}
	service := application.NewAuthService(repo)

	userID, err := service.ValidateToken(context.Background(), "valid_token")

	assert.NoError(t, err)
	assert.Equal(t, "user123", userID)

	_, err = service.ValidateToken(context.Background(), "invalid_token")
	assert.Error(t, err)
}

type MockAuthRepository struct {
	ValidToken string
	UserID     string
}

func (m *MockAuthRepository) ValidateToken(token string) (string, error) {
	if token == m.ValidToken {
		return m.UserID, nil
	}
	return "", domain.ErrInvalidToken
}
