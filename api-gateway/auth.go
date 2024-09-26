package main

import (
	"context"

	authpb "github.com/hossein-225/Library-Management/auth-service/proto"
	"google.golang.org/grpc"
)

func authenticateUser(ctx context.Context, token string) (string, bool, error) {
	conn, err := grpc.NewClient("auth-service:50054", grpc.WithInsecure())
	if err != nil {
		return "", false, err
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)
	req := &authpb.ValidateTokenRequest{Token: token}
	res, err := client.ValidateToken(ctx, req)
	if err != nil {
		return "", false, err
	}

	return res.UserId, res.Role == authpb.Role_ADMIN, nil
}
