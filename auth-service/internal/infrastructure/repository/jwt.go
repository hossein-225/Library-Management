package repository

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hossein-225/Library-Management/auth-service/pkg/utils"
)

type JWTAuthRepository struct{}

func NewJWTAuthRepository() *JWTAuthRepository {
	return &JWTAuthRepository{}
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
