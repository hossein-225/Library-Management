package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	_ "github.com/hossein-225/Library-Management/book-service/docs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/hossein-225/Library-Management/book-service/internal/application"
	"github.com/hossein-225/Library-Management/book-service/internal/domain"
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
	db := configDB()
	defer db.Close()

	log.Println("connect to postgresql successfully")

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

func configDB() *sql.DB {
	client, err := gorm.Open(postgres.Open("postgres://"+os.Getenv("PG_USER")+
		":"+os.Getenv("PG_PASSWORD")+"@"+os.Getenv("PG_URL")+":"+
		os.Getenv("PG_PORT")+"/"+os.Getenv("PG_NAME")), &gorm.Config{})
	if err != nil {
		log.Println("couldn't connect to postgresql DB", err)
		log.Fatal(err)
	}

	var sqlDB *sql.DB
	sqlDB, err = client.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = ConfigModels(client)
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return sqlDB
}

func ConfigModels(client *gorm.DB) error {

	err := client.AutoMigrate(&domain.Book{})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Table Created")

	return nil

}
