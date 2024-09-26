package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// _ "path/to/your/project/docs"
)

// @title Library Management API
// @version 1.0
// @description API documentation for the Library Management system

// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/books", handleBookList)
	router.POST("/books", handleAddBook)
	router.PUT("/books/:id", handleUpdateBook)
	router.DELETE("/books/:id", handleDeleteBook)

	router.POST("/register", handleUserRegister)
	router.POST("/login", handleUserLogin)

	router.POST("/borrow", handleBorrowBook)
	router.POST("/return", handleReturnBook)

	log.Println("API Gateway running on port 8080")
	router.Run(":8080")
}
