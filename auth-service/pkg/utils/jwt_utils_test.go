package utils_test

import (
	"testing"

	"github.com/hossein-225/Library-Management/auth-service/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	token, err := utils.GenerateJWT("user123", "admin")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := utils.ParseJWT(token)
	assert.NoError(t, err)
	assert.Equal(t, "user123", claims["userID"])
}

func TestParseJWT_InvalidToken(t *testing.T) {
	_, err := utils.ParseJWT("invalid_token")
	assert.Error(t, err)
}
