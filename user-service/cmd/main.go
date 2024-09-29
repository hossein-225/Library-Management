package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hossein-225/Library-Management/user-service/internal/application"
	"github.com/hossein-225/Library-Management/user-service/internal/domain"
	user_grpc "github.com/hossein-225/Library-Management/user-service/internal/infrastructure/grpc"
	"github.com/hossein-225/Library-Management/user-service/internal/infrastructure/repository"
	pb "github.com/hossein-225/Library-Management/user-service/proto"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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
	db := configDB()
	defer db.Close()

	log.Println("connect to postgresql successfully")

	repo := repository.NewPostgresUserRepository(db)
	service := application.NewUserService(repo)
	grpcServer := user_grpc.NewUserGRPCServer(service)

	createAdminIfNotExists(db)

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

	err = configModels(client)
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return sqlDB
}

func createAdminIfNotExists(db *sql.DB) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = 'admin@example.com'").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to check for admin account: %v", err)
	}

	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash admin password: %v", err)
		}

		_, err = db.Exec("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)", "admin", "admin@example.com", string(hashedPassword), "admin")
		if err != nil {
			log.Fatalf("Failed to create admin user: %v", err)
		}

		log.Println("Admin account created successfully!")
	} else {
		log.Println("Admin account already exists.")
	}
}

func configModels(client *gorm.DB) error {

	err := client.Migrator().AutoMigrate(&domain.User{})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Table Created")

	return nil

}
