package application_test

import (
	"context"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hossein-225/Library-Management/auth-service/internal/application"
	"github.com/hossein-225/Library-Management/auth-service/internal/domain"
	"github.com/hossein-225/Library-Management/auth-service/proto"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	repo := &MockAuthRepository{
		GeneratedToken: "mocked_valid_token",
	}
	service := application.NewAuthService(repo)

	token, err := service.GenerateToken(context.Background(), "user123", proto.Role_USER)

	assert.NoError(t, err)
	assert.Equal(t, "mocked_valid_token", token)
}

func TestValidateToken(t *testing.T) {
	repo := &MockAuthRepository{
		ValidToken: "valid_token",
		UserID:     "user123",
		Role:       "user",
	}
	service := application.NewAuthService(repo)

	claims, err := service.ValidateToken(context.Background(), "valid_token")
	assert.NoError(t, err)

	userID, ok := claims["userID"].(string)
	assert.True(t, ok)
	assert.Equal(t, "user123", userID)

	role, ok := claims["role"].(string)
	assert.True(t, ok)
	assert.Equal(t, "user", role)

	_, err = service.ValidateToken(context.Background(), "invalid_token")
	assert.Error(t, err)
}

type MockAuthRepository struct {
	GeneratedToken string
	ValidToken     string
	UserID         string
	Role           string
}

func (m *MockAuthRepository) GenerateToken(userID string, role proto.Role) (string, error) {
	return m.GeneratedToken, nil
}

func (m *MockAuthRepository) ValidateToken(token string) (jwt.MapClaims, error) {
	if token == m.ValidToken {
		claims := jwt.MapClaims{
			"userID": m.UserID,
			"role":   m.Role,
		}
		return claims, nil
	}
	return nil, domain.ErrInvalidToken
}
