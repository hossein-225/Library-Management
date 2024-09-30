package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	borrowpb "github.com/hossein-225/Library-Management/borrow-service/proto"
	borrowpbMock "github.com/hossein-225/Library-Management/borrow-service/proto/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandleBorrowBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBorrowClient := borrowpbMock.NewMockBorrowServiceClient(ctrl)

	mockBorrowClient.EXPECT().BorrowBook(gomock.Any(), gomock.Any()).Return(&borrowpb.BorrowBookResponse{}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/books/borrow", func(c *gin.Context) {
		bookID := c.PostForm("book_id")
		userID := "user123"

		err := borrowBook(c.Request.Context(), mockBorrowClient, userID, bookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to borrow book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book borrowed successfully"})
	})

	reqBody := `book_id=book123`
	req, _ := http.NewRequest("POST", "/books/borrow", strings.NewReader(reqBody))
	req.Header.Set("Authorization", "Bearer valid_token")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book borrowed successfully")
}

func TestHandleReturnBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBorrowClient := borrowpbMock.NewMockBorrowServiceClient(ctrl)

	mockBorrowClient.EXPECT().ReturnBook(gomock.Any(), gomock.Any()).Return(&borrowpb.ReturnBookResponse{}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/books/return", func(c *gin.Context) {
		bookID := c.PostForm("book_id")
		userID := "user123"

		err := returnBook(c.Request.Context(), mockBorrowClient, userID, bookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to return book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
	})

	reqBody := `book_id=book123`
	req, _ := http.NewRequest("POST", "/books/return", strings.NewReader(reqBody))
	req.Header.Set("Authorization", "Bearer valid_token")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book returned successfully")
}
