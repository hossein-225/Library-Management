package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleBorrowBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "valid_token" {
			c.Set("userID", "user123")
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
	})

	router.POST("/borrow", func(c *gin.Context) {
		var request struct {
			BookID string `json:"book_id"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book borrowed successfully"})
	})

	reqBody := `{"book_id": "book123"}`
	req, _ := http.NewRequest("POST", "/borrow", strings.NewReader(reqBody))
	req.Header.Set("Authorization", "valid_token")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book borrowed successfully")
}

func TestHandleReturnBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "valid_token" {
			c.Set("userID", "user123")
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
	})

	router.POST("/return", func(c *gin.Context) {
		var request struct {
			BookID string `json:"book_id"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
	})

	reqBody := `{"book_id": "book123"}`
	req, _ := http.NewRequest("POST", "/return", strings.NewReader(reqBody))
	req.Header.Set("Authorization", "valid_token")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book returned successfully")
}
