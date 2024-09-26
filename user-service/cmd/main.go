package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hossein-225/Library-Management/user-service/internal/application"
	user_grpc "github.com/hossein-225/Library-Management/user-service/internal/infrastructure/grpc"
	"github.com/hossein-225/Library-Management/user-service/internal/infrastructure/repository"
	pb "github.com/hossein-225/Library-Management/user-service/proto"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net"

	_ "github.com/hossein-225/Library-Management/user-service/docs"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

// @title Library Management API - user-service
// @version 0.0.6
// @description API documentation for the Library Management system - user-service

// @host user-service:50052
// @BasePath /
func main() {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=user_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewPostgresUserRepository(db)
	service := application.NewUserService(repo)
	grpcServer := user_grpc.NewUserGRPCServer(service)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, grpcServer)

	go func() {
		log.Println("User Service is running on port 50052...")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Swagger is available at http://localhost:8080/swagger/index.html")
	router.Run(":8080")
}
