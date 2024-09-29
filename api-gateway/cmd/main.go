package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/hossein-225/Library-Management/api-gateway/docs"
	"github.com/hossein-225/Library-Management/api-gateway/handler"
	cors "github.com/hossein-225/Library-Management/api-gateway/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Library Management API
// @version 0.0.6
// @description API documentation for the Library Management system

// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.CORSMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{

		books := v1.Group("/books")
		{
			books.GET("", handler.HandleBookList)
			books.POST("", handler.HandleAddBook)
			books.PUT("/:id", handler.HandleUpdateBook)
			books.DELETE("/:id", handler.HandleDeleteBook)

			books.POST("/borrow", handler.HandleBorrowBook)
			books.POST("/return", handler.HandleReturnBook)
		}

		users := v1.Group("/users")
		{
			users.POST("/register", handler.HandleUserRegister)
			users.POST("/login", handler.HandleUserLogin)

			users.GET("/profile", handler.HandleGetUserProfile)
			users.PUT("/profile", handler.HandleUpdateUserProfile)
		}

	}

	log.Println("API Gateway running on port 8080")
	router.Run(":8080")
}
