package application

import (
	"context"

	"github.com/hossein-225/Library-Management/auth-service/internal/domain"
	pb "github.com/hossein-225/Library-Management/auth-service/proto"
)

type AuthService struct {
	repo domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(ctx context.Context, userID string, role pb.Role) (string, error) {
	return s.repo.GenerateToken(userID, role)
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (string, error) {
	return s.repo.ValidateToken(token)
}
