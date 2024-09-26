package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	_ "github.com/hossein-225/Library-Management/borrow-service/docs"
	"github.com/hossein-225/Library-Management/borrow-service/internal/application"
	borrow_grpc "github.com/hossein-225/Library-Management/borrow-service/internal/infrastructure/grpc"
	"github.com/hossein-225/Library-Management/borrow-service/internal/infrastructure/repository"
	pb "github.com/hossein-225/Library-Management/borrow-service/proto"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Library Management API - borrow-service
// @version 0.0.6
// @description API documentation for the Library Management system - borrow-service

// @host borrow-service:50053
// @BasePath /
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

	go func() {
		log.Println("Borrow Service is running on port 50053...")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Swagger is available at http://localhost:8080/swagger/index.html")
	router.Run(":8080")
}
