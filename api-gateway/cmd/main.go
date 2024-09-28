package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/hossein-225/Library-Management/api-gateway/docs"
	"github.com/hossein-225/Library-Management/api-gateway/handler"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/books", handler.HandleBookList)
	router.POST("/books", handler.HandleAddBook)
	router.PUT("/books/:id", handler.HandleUpdateBook)
	router.DELETE("/books/:id", handler.HandleDeleteBook)

	router.POST("/register", handler.HandleUserRegister)
	router.POST("/login", handler.HandleUserLogin)

	router.POST("/borrow", handler.HandleBorrowBook)
	router.POST("/return", handler.HandleReturnBook)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("API Gateway running on port 8080")
	router.Run(":8080")
}
