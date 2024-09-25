package grpc

import (
	"context"

	"github.com/hossein-225/Library-Management/user-service/internal/application"
	"github.com/hossein-225/Library-Management/user-service/internal/domain"
	"github.com/hossein-225/Library-Management/user-service/pkg/utils"
	pb "github.com/hossein-225/Library-Management/user-service/proto"
)

type UserGRPCServer struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserGRPCServer(service *application.UserService) *UserGRPCServer {
	return &UserGRPCServer{service: service}
}

func (s *UserGRPCServer) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := &domain.User{
		ID:       utils.GenerateUUID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.service.RegisterUser(ctx, user); err != nil {
		return nil, err
	}

	return &pb.RegisterUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (s *UserGRPCServer) AuthenticateUser(ctx context.Context, req *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	user, err := s.service.AuthenticateUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticateUserResponse{
		Token: token,
	}, nil
}

func (s *UserGRPCServer) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	user, err := s.service.GetUserProfile(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserProfileResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (s *UserGRPCServer) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	user := &domain.User{
		ID:    req.UserId,
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.service.UpdateUserProfile(ctx, user); err != nil {
		return nil, err
	}

	return &pb.UpdateUserProfileResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
