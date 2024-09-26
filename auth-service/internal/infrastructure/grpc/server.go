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
// @Summary Generate a JWT token for a user with a role
// @Description Generates a JWT token for the user with the provided user ID and role
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user_id  body   string   true  "User ID"
// @Param   role     body   string   true  "User role"
// @Success 200 {object} map[string]interface{} "Token generated successfully"
// @Failure 400 {string} string "User ID or Role cannot be empty"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/token [post]
func (s *AuthGRPCServer) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	if req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "User ID cannot be empty")
	}

	token, err := s.service.GenerateToken(ctx, req.UserId, req.Role)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}

	return &pb.GenerateTokenResponse{Token: token}, nil
}

// ValidateToken
// @Summary Validate a JWT token
// @Description Validates the provided JWT token and returns the associated user ID and role
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   token  body   string   true  "JWT token"
// @Success 200 {object} map[string]interface{} "Token validated successfully"
// @Failure 400 {string} string "Token cannot be empty or is invalid"
// @Failure 401 {string} string "Invalid token: missing userID or role"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/validate [post]
func (s *AuthGRPCServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	if req.Token == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Token cannot be empty")
	}

	claims, err := s.service.ValidateToken(ctx, req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to validate token: %v", err)
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token: missing userID")
	}

	role, ok := claims["role"].(pb.Role)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token: missing role")
	}

	return &pb.ValidateTokenResponse{UserId: userID, Role: role}, nil
}
