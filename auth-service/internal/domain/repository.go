package domain

import (
	pb "github.com/hossein-225/Library-Management/auth-service/proto"
)

type AuthRepository interface {
	GenerateToken(userID string, role pb.Role) (string, error)
	ValidateToken(token string) (string, error)
}
