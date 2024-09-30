package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/hossein-225/Library-Management/user-service/internal/application"
	"github.com/hossein-225/Library-Management/user-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) RegisterUser(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) AuthenticateUser(email, password string) (*domain.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetUserProfile(email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserProfile(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestRegisterUser(t *testing.T) {
	repo := new(MockUserRepository)
	service := application.NewUserService(repo)

	user := &domain.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		Role:     "user",
	}

	repo.On("RegisterUser", user).Return(nil)

	err := service.RegisterUser(context.Background(), user)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestAuthenticateUser(t *testing.T) {
	repo := new(MockUserRepository)
	service := application.NewUserService(repo)

	user := &domain.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		Role:     "user",
	}

	repo.On("AuthenticateUser", "test@example.com", "password").Return(user, nil)

	result, err := service.AuthenticateUser(context.Background(), "test@example.com", "password")

	assert.NoError(t, err)
	assert.Equal(t, "Test User", result.Name)
	repo.AssertExpectations(t)
}

func TestGetUserProfile(t *testing.T) {
	repo := new(MockUserRepository)
	service := application.NewUserService(repo)

	user := &domain.User{
		Name:  "Test User",
		Email: "test@example.com",
		Role:  "user",
	}

	repo.On("GetUserProfile", "test@example.com").Return(user, nil)

	result, err := service.GetUserProfile(context.Background(), "test@example.com")

	assert.NoError(t, err)
	assert.Equal(t, "Test User", result.Name)
	assert.Equal(t, "user", result.Role)
	repo.AssertExpectations(t)
}

func TestUpdateUserProfile(t *testing.T) {
	repo := new(MockUserRepository)
	service := application.NewUserService(repo)

	user := &domain.User{
		Name:  "Updated User",
		Email: "test@example.com",
		Role:  "user",
	}

	repo.On("UpdateUserProfile", user).Return(nil)

	err := service.UpdateUserProfile(context.Background(), user)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUpdateUserProfile_Error(t *testing.T) {
	repo := new(MockUserRepository)
	service := application.NewUserService(repo)

	user := &domain.User{
		Name:  "Updated User",
		Email: "test@example.com",
		Role:  "user",
	}

	repo.On("UpdateUserProfile", user).Return(errors.New("update error"))

	err := service.UpdateUserProfile(context.Background(), user)

	assert.Error(t, err)
	assert.Equal(t, "update error", err.Error())
	repo.AssertExpectations(t)
}
