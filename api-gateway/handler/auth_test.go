package handler

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	authpb "github.com/hossein-225/Library-Management/auth-service/proto"
	mock_proto "github.com/hossein-225/Library-Management/auth-service/proto/mock"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthServiceClient := mock_proto.NewMockAuthServiceClient(ctrl)

	validToken := "valid_token"
	userID := "123"
	isAdmin := true

	mockAuthServiceClient.
		EXPECT().
		ValidateToken(gomock.Any(), &authpb.ValidateTokenRequest{Token: validToken}).
		Return(&authpb.ValidateTokenResponse{
			UserId: userID,
			Role:   authpb.Role_ADMIN,
		}, nil)

	conn := mockAuthServiceClient

	userId, isAdminResult, err := authenticateUser(context.Background(), conn, validToken)

	assert.NoError(t, err)
	assert.Equal(t, userID, userId)
	assert.Equal(t, isAdmin, isAdminResult)
}
