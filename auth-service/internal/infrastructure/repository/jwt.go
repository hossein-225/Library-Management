package repository

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hossein-225/Library-Management/auth-service/pkg/utils"
	pb "github.com/hossein-225/Library-Management/auth-service/proto"
)

type JWTAuthRepository struct{}

func NewJWTAuthRepository() *JWTAuthRepository {
	return &JWTAuthRepository{}
}

func (r *JWTAuthRepository) GenerateToken(userID string, role pb.Role) (string, error) {
	var strRole string

	if role == pb.Role_ADMIN {
		strRole = "admin"
	} else if role == pb.Role_USER {
		strRole = "user"
	}

	return utils.GenerateJWT(userID, strRole)
}

func (r *JWTAuthRepository) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return utils.JwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["userID"].(string)
		return userID, nil
	} else {
		return "", errors.New("invalid token")
	}
}
