package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/hossein-225/Library-Management/borrow-service/internal/application"
	borrow_grpc "github.com/hossein-225/Library-Management/borrow-service/internal/infrastructure/grpc"
	"github.com/hossein-225/Library-Management/borrow-service/internal/infrastructure/repository"
	pb "github.com/hossein-225/Library-Management/borrow-service/proto"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=borrow_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewPostgresBorrowRepository(db)
	service := application.NewBorrowService(repo)
	grpcServer := borrow_grpc.NewBorrowGRPCServer(service)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBorrowServiceServer(s, grpcServer)

	log.Println("Borrow Service is running on port 50053...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
