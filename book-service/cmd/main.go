package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/hossein-225/Library-Management/book-service/internal/application"
	book_grpc "github.com/hossein-225/Library-Management/book-service/internal/infrastructure/grpc"
	"github.com/hossein-225/Library-Management/book-service/internal/infrastructure/repository"
	pb "github.com/hossein-225/Library-Management/book-service/proto"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=book_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewPostgresBookRepository(db)
	service := application.NewBookService(repo)
	grpcServer := book_grpc.NewBookGRPCServer(service)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, grpcServer)

	log.Println("Book Service is running on port 50051...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
