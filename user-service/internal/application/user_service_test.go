package application_test

import (
	"context"
	"testing"

	"github.com/hossein-225/Library-Management/user-service/internal/application"
	"github.com/hossein-225/Library-Management/user-service/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	repo := &MockUserRepository{}
	service := application.NewUserService(repo)

	user := &domain.User{
		ID:       "123",
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		Role:     "user",
	}

	err := service.RegisterUser(context.Background(), user)

	assert.NoError(t, err)
	assert.Equal(t, user, repo.RegisteredUser)
}

func TestAuthenticateUser(t *testing.T) {
	repo := &MockUserRepository{
		Users: []*domain.User{
			{
				ID:       "123",
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password",
				Role:     "user",
			},
		},
	}
	service := application.NewUserService(repo)

	user, err := service.AuthenticateUser(context.Background(), "test@example.com", "password")

	assert.NoError(t, err)
	assert.Equal(t, "Test User", user.Name)
}

type MockUserRepository struct {
	Users          []*domain.User
	RegisteredUser *domain.User
}

func (m *MockUserRepository) RegisterUser(user *domain.User) error {
	m.RegisteredUser = user
	return nil
}

func (m *MockUserRepository) AuthenticateUser(email, password string) (*domain.User, error) {
	for _, user := range m.Users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return nil, nil
}

func (m *MockUserRepository) GetUserProfile(id string) (*domain.User, error) {
	return nil, nil
}

func (m *MockUserRepository) UpdateUserProfile(user *domain.User) error {
	return nil
}
