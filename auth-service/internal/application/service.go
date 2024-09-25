package application

import (
	"context"

	"github.com/hossein-225/Library-Management/auth-service/internal/domain"
	"github.com/hossein-225/Library-Management/auth-service/pkg/utils"
)

type AuthService struct {
	repo domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(ctx context.Context, userID string) (string, error) {
	return utils.GenerateJWT(userID)
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (string, error) {
	return s.repo.ValidateToken(token)
}
