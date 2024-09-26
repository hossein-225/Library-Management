package grpc

import (
	"context"

	"github.com/hossein-225/Library-Management/auth-service/internal/application"
	pb "github.com/hossein-225/Library-Management/auth-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthGRPCServer struct {
	pb.UnimplementedAuthServiceServer
	service *application.AuthService
}

func NewAuthGRPCServer(service *application.AuthService) *AuthGRPCServer {
	return &AuthGRPCServer{service: service}
}

// GenerateToken
// @Summary Generate a JWT token for a user
// @Description Generates a JWT token for the user with the provided user ID
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user_id  body   string   true  "User ID"
// @Success 200 {object} pb.GenerateTokenResponse "Token generated successfully"
// @Failure 400 {string} string "User ID cannot be empty"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/token [post]
func (s *AuthGRPCServer) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	if req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "User ID cannot be empty")
	}

	token, err := s.service.GenerateToken(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}

	return &pb.GenerateTokenResponse{Token: token}, nil
}

// ValidateToken
// @Summary Validate a JWT token
// @Description Validates the provided JWT token and returns the associated user ID
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   token  body   string   true  "JWT token"
// @Success 200 {object} pb.ValidateTokenResponse "Token validated successfully"
// @Failure 400 {string} string "Token cannot be empty or is invalid"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/validate [post]
func (s *AuthGRPCServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	if req.Token == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Token cannot be empty")
	}

	userID, err := s.service.ValidateToken(ctx, req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to validate token: %v", err)
	}

	return &pb.ValidateTokenResponse{UserId: userID}, nil
}
