package grpc

import (
	"context"

	"github.com/hossein-225/Library-Management/user-service/internal/application"
	"github.com/hossein-225/Library-Management/user-service/internal/domain"
	"github.com/hossein-225/Library-Management/user-service/pkg/utils"
	pb "github.com/hossein-225/Library-Management/user-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserGRPCServer struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserGRPCServer(service *application.UserService) *UserGRPCServer {
	return &UserGRPCServer{service: service}
}

// RegisterUser
// @Summary Register a new user
// @Description Registers a new user with the provided information
// @Tags users
// @Accept  json
// @Produce  json
// @Param   name     body   string   true  "Name of the user"
// @Param   email    body   string   true  "Email of the user"
// @Param   password body   string   true  "Password of the user"
// @Success 200 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {string} string "Name or email cannot be empty"
// @Failure 500 {string} string "Failed to register user"
// @Router /users/register [post]
func (s *UserGRPCServer) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	if req.Name == "" || req.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name or email cannot be empty")
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     "user",
	}

	if err := s.service.RegisterUser(ctx, user); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to register user: %v", err)
	}

	return &pb.RegisterUserResponse{
		User: &pb.User{
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}

// AuthenticateUser
// @Summary Authenticate a user
// @Description Authenticate a user and return a JWT token
// @Tags users
// @Accept  json
// @Produce  json
// @Param   email    body   string   true  "Email of the user"
// @Param   password body   string   true  "Password of the user"
// @Success 200 {object} map[string]interface{} "User authenticated successfully"
// @Failure 400 {string} string "Invalid email or password"
// @Failure 500 {string} string "Failed to generate token"
// @Router /users/authenticate [post]
func (s *UserGRPCServer) AuthenticateUser(ctx context.Context, req *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email or password cannot be empty")
	}

	user, err := s.service.AuthenticateUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password")
	}

	token, err := utils.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token")
	}

	return &pb.AuthenticateUserResponse{
		Token: token,
	}, nil
}

// GetUserProfile
// @Summary Get a user's profile
// @Description Retrieves the profile information for the authenticated user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   email  query  string   true  "Email"
// @Success 200 {object} map[string]interface{} "Profile retrieved successfully"
// @Failure 400 {string} string "Email cannot be empty"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users/profile [get]
func (s *UserGRPCServer) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	if req.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email cannot be empty")
	}

	user, err := s.service.GetUserProfile(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve user profile: %v", err)
	}

	return &pb.GetUserProfileResponse{
		User: &pb.User{
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

// UpdateUserProfile
// @Summary Update a user's profile
// @Description Update profile information for the authenticated user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   email  body   string   true  "Email"
// @Param   name     body   string   true  "New name"
// @Param   email    body   string   true  "New email"
// @Success 200 {object} map[string]interface{} "Profile updated successfully"
// @Failure 400 {string} string "Email, name, or email cannot be empty"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users/profile [put]
func (s *UserGRPCServer) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	if req.Name == "" || req.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email, name, or email cannot be empty")
	}

	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.service.UpdateUserProfile(ctx, user); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user profile: %v", err)
	}

	return &pb.UpdateUserProfileResponse{
		User: &pb.User{
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
