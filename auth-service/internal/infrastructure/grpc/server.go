package grpc

import (
	"context"

	"github.com/hossein-225/Library-Management/auth-service/internal/application"
	pb "github.com/hossein-225/Library-Management/auth-service/proto"
)

type AuthGRPCServer struct {
	pb.UnimplementedAuthServiceServer
	service *application.AuthService
}

func NewAuthGRPCServer(service *application.AuthService) *AuthGRPCServer {
	return &AuthGRPCServer{service: service}
}

func (s *AuthGRPCServer) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	token, err := s.service.GenerateToken(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GenerateTokenResponse{Token: token}, nil
}

func (s *AuthGRPCServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	userID, err := s.service.ValidateToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.ValidateTokenResponse{UserId: userID}, nil
}
