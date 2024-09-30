package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	bookpb "github.com/hossein-225/Library-Management/book-service/proto"
	mock_proto "github.com/hossein-225/Library-Management/book-service/proto/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandleBookList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookClient := mock_proto.NewMockBookServiceClient(ctrl)

	mockBookClient.EXPECT().ListBooks(gomock.Any(), gomock.Any()).Return(&bookpb.ListBooksResponse{
		Books: []*bookpb.Book{
			{
				Id:       "1",
				Title:    "Test Book",
				Author:   "John Doe",
				Category: "Science",
			},
		},
	}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/books", func(c *gin.Context) {
		books, err := fetchBooks(c.Request.Context(), mockBookClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
			return
		}
		c.JSON(http.StatusOK, books)
	})

	req, _ := http.NewRequest("GET", "/books", nil)
	req.Header.Set("Authorization", "valid_token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Book")
}

func TestHandleAddBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookClient := mock_proto.NewMockBookServiceClient(ctrl)

	mockBookClient.EXPECT().AddBook(gomock.Any(), gomock.Any()).Return(&bookpb.AddBookResponse{}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/books", func(c *gin.Context) {
		title := c.PostForm("title")
		author := c.PostForm("author")
		category := c.PostForm("category")

		err := addBook(c.Request.Context(), mockBookClient, title, author, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookClient := mock_proto.NewMockBookServiceClient(ctrl)

	mockBookClient.EXPECT().UpdateBook(gomock.Any(), gomock.Any()).Return(&bookpb.UpdateBookResponse{}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.PUT("/books/:id", func(c *gin.Context) {
		bookID := c.Param("id")
		title := c.PostForm("title")
		author := c.PostForm("author")
		category := c.PostForm("category")

		err := updateBook(c.Request.Context(), mockBookClient, bookID, title, author, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookClient := mock_proto.NewMockBookServiceClient(ctrl)

	mockBookClient.EXPECT().DeleteBook(gomock.Any(), gomock.Any()).Return(&bookpb.DeleteBookResponse{}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.DELETE("/books/:id", func(c *gin.Context) {
		bookID := c.Param("id")

		err := deleteBook(c.Request.Context(), mockBookClient, bookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
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

func TestHandleSearchBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookClient := mock_proto.NewMockBookServiceClient(ctrl)

	mockBookClient.EXPECT().SearchBooks(gomock.Any(), gomock.Any()).Return(&bookpb.SearchBooksResponse{
		Books: []*bookpb.Book{
			{
				Id:       "1",
				Title:    "Test Book",
				Author:   "John Doe",
				Category: "Science",
			},
		},
	}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/books/search", func(c *gin.Context) {
		title := c.Query("title")
		author := c.Query("author")
		category := c.Query("category")

		books, err := searchBooks(c.Request.Context(), mockBookClient, title, author, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search books"})
			return
		}
		c.JSON(http.StatusOK, books)
	})

	req, _ := http.NewRequest("GET", "/books/search?title=Test Book", nil)
	req.Header.Set("Authorization", "valid_token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Book")
}
