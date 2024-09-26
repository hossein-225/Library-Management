package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/hossein-225/Library-Management/book-service/docs"

	"github.com/gin-gonic/gin"
	"github.com/hossein-225/Library-Management/book-service/internal/application"
	book_grpc "github.com/hossein-225/Library-Management/book-service/internal/infrastructure/grpc"
	"github.com/hossein-225/Library-Management/book-service/internal/infrastructure/repository"
	pb "github.com/hossein-225/Library-Management/book-service/proto"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Library Management API - book-service
// @version 0.0.6
// @description API documentation for the Library Management system - book-service

// @host book-service:50051
// @BasePath /
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

	go func() {
		log.Println("Book Service is running on port 50051...")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Swagger is available at http://localhost:8080/swagger/index.html")
	router.Run(":8080")
}
