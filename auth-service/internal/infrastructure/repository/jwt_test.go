package repository_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hossein-225/Library-Management/auth-service/internal/infrastructure/repository"
	"github.com/hossein-225/Library-Management/auth-service/pkg/utils"
	pb "github.com/hossein-225/Library-Management/auth-service/proto"
	"github.com/stretchr/testify/assert"
)

func init() {
	utils.JwtSecret = []byte("mock_secret")
}

func TestJWTAuthRepository_GenerateToken(t *testing.T) {
	repo := repository.NewJWTAuthRepository()

	token, err := repo.GenerateToken("user123", pb.Role_USER)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return utils.JwtSecret, nil
	})
	assert.NoError(t, err)
	assert.NotNil(t, parsedToken)

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		assert.Equal(t, "user123", claims["userID"])
		assert.Equal(t, "user", claims["role"])
	} else {
		t.FailNow()
	}
}

func TestJWTAuthRepository_GenerateToken_AdminRole(t *testing.T) {
	repo := repository.NewJWTAuthRepository()

	token, err := repo.GenerateToken("admin123", pb.Role_ADMIN)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return utils.JwtSecret, nil
	})
	assert.NoError(t, err)
	assert.NotNil(t, parsedToken)

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		assert.Equal(t, "admin123", claims["userID"])
		assert.Equal(t, "admin", claims["role"])
	} else {
		t.FailNow()
	}
}

func TestJWTAuthRepository_ValidateToken(t *testing.T) {
	repo := repository.NewJWTAuthRepository()

	token, err := utils.GenerateJWT("user123", "user")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := repo.ValidateToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, claims)

	assert.Equal(t, "user123", claims["userID"])
	assert.Equal(t, "user", claims["role"])
}

func TestJWTAuthRepository_ValidateToken_InvalidToken(t *testing.T) {
	repo := repository.NewJWTAuthRepository()

	invalidToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": "user123",
		"role":   "user",
		"exp":    time.Now().Add(-time.Hour).Unix(),
	})
	tokenString, err := invalidToken.SignedString(utils.JwtSecret)
	assert.NoError(t, err)

	_, err = repo.ValidateToken(tokenString)
	assert.Error(t, err)
	assert.Equal(t, "Token is expired", err.Error())
}
