package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	_ "github.com/hossein-225/Library-Management/auth-service/docs"
	"github.com/hossein-225/Library-Management/auth-service/internal/application"
	auth_grpc "github.com/hossein-225/Library-Management/auth-service/internal/infrastructure/grpc"
	"github.com/hossein-225/Library-Management/auth-service/internal/infrastructure/repository"
	pb "github.com/hossein-225/Library-Management/auth-service/proto"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func main() {
	repo := repository.NewJWTAuthRepository()
	service := application.NewAuthService(repo)
	grpcServer := auth_grpc.NewAuthGRPCServer(service)

	listener, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, grpcServer)

	go func() {
		log.Println("Auth Service is running on port 50054...")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Swagger is available at http://localhost:8080/swagger/index.html")
	router.Run(":8080")
}
