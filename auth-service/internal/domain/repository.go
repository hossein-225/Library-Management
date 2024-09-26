package domain

import (
	"github.com/golang-jwt/jwt/v4"
	pb "github.com/hossein-225/Library-Management/auth-service/proto"
)

type AuthRepository interface {
	GenerateToken(userID string, role pb.Role) (string, error)
	ValidateToken(token string) (jwt.MapClaims, error)
}
