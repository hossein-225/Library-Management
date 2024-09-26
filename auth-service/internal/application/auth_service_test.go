package application_test

import (
	"context"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hossein-225/Library-Management/auth-service/internal/application"
	"github.com/hossein-225/Library-Management/auth-service/internal/domain"
	"github.com/hossein-225/Library-Management/auth-service/pkg/utils"

	"github.com/hossein-225/Library-Management/auth-service/proto"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	repo := &MockAuthRepository{
		ValidToken: "valid_token",
		UserID:     "user123",
	}
	service := application.NewAuthService(repo)

	token, err := service.GenerateToken(context.Background(), "user123", 1)

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

	claims, err := service.ValidateToken(context.Background(), "valid_token")

	assert.NoError(t, err)

	userID, ok := claims["userID"].(string)
	assert.True(t, ok)

	assert.Equal(t, "user123", userID)

	_, err = service.ValidateToken(context.Background(), "invalid_token")
	assert.Error(t, err)
}

type MockAuthRepository struct {
	ValidToken string
	UserID     string
}

func (m *MockAuthRepository) GenerateToken(userID string, role proto.Role) (string, error) {
	var strRole string

	if role == proto.Role_ADMIN {
		strRole = "admin"
	} else if role == proto.Role_USER {
		strRole = "user"
	}
	return utils.GenerateJWT(userID, strRole)
}

func (m *MockAuthRepository) ValidateToken(token string) (jwt.MapClaims, error) {
	if token == m.ValidToken {
		claims := jwt.MapClaims{
			"userID": m.UserID,
		}
		return claims, nil
	}
	return nil, domain.ErrInvalidToken
}
