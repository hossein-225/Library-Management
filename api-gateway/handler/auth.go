package handler

import (
	"context"

	authpb "github.com/hossein-225/Library-Management/auth-service/proto"
)

func authenticateUser(ctx context.Context, client authpb.AuthServiceClient, token string) (string, bool, error) {
	req := &authpb.ValidateTokenRequest{Token: token}
	res, err := client.ValidateToken(ctx, req)
	if err != nil {
		return "", false, err
	}

	return res.UserId, res.Role == authpb.Role_ADMIN, nil
}
