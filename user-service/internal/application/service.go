package application

import (
	"context"

	"github.com/hossein-225/Library-Management/user-service/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, user *domain.User) error {
	return s.repo.RegisterUser(user)
}

func (s *UserService) AuthenticateUser(ctx context.Context, email, password string) (*domain.User, error) {
	return s.repo.AuthenticateUser(email, password)
}

func (s *UserService) GetUserProfile(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetUserProfile(id)
}

func (s *UserService) UpdateUserProfile(ctx context.Context, user *domain.User) error {
	return s.repo.UpdateUserProfile(user)
}
