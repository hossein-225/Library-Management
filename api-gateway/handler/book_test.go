package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleBookList(t *testing.T) {
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

	router.GET("/books", func(c *gin.Context) {
		books := []string{"Book List"}
		c.JSON(http.StatusOK, books)
	})

	req, _ := http.NewRequest("GET", "/books", nil)
	req.Header.Set("Authorization", "valid_token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book List")
}

func TestHandleAddBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "valid_admin_token" {
			c.Set("userID", "admin123")
			c.Set("isAdmin", true)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
	})

	router.POST("/books", func(c *gin.Context) {
		var book struct {
			Title    string `json:"title"`
			Author   string `json:"author"`
			Category string `json:"category"`
		}

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})
	})

	reqBody := `{"title": "New Book", "author": "John Doe", "category": "Science"}`
	req, _ := http.NewRequest("POST", "/books", strings.NewReader(reqBody))
	req.Header.Set("Authorization", "valid_admin_token")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book added successfully")
}

func TestHandleUpdateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "valid_admin_token" {
			c.Set("userID", "admin123")
			c.Set("isAdmin", true)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
	})

	router.PUT("/books/:id", func(c *gin.Context) {
		var book struct {
			Title    string `json:"title"`
			Author   string `json:"author"`
			Category string `json:"category"`
		}

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
	})

	reqBody := `{"title": "Updated Book", "author": "Jane Doe", "category": "Fiction"}`
	req, _ := http.NewRequest("PUT", "/books/book123", strings.NewReader(reqBody))
	req.Header.Set("Authorization", "valid_admin_token")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book updated successfully")
}

func TestHandleDeleteBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "valid_admin_token" {
			c.Set("userID", "admin123")
			c.Set("isAdmin", true)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
	})

	router.DELETE("/books/:id", func(c *gin.Context) {
		bookID := c.Param("id")
		if bookID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	})

	req, _ := http.NewRequest("DELETE", "/books/book123", nil)
	req.Header.Set("Authorization", "valid_admin_token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book deleted successfully")
}
