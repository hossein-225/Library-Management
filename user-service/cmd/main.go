package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/hossein-225/Library-Management/user-service/internal/application"
	user_grpc "github.com/hossein-225/Library-Management/user-service/internal/infrastructure/grpc"
	"github.com/hossein-225/Library-Management/user-service/internal/infrastructure/repository"
	pb "github.com/hossein-225/Library-Management/user-service/proto"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

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

	log.Println("User Service is running on port 50052...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
