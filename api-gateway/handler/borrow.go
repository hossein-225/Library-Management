package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	authpb "github.com/hossein-225/Library-Management/auth-service/proto"
	borrowpb "github.com/hossein-225/Library-Management/borrow-service/proto"
	"google.golang.org/grpc"
)

// @Summary Borrow a book
// @Description Borrow a book from the library
// @Tags borrow
// @Accept json
// @Produce json
// @Param book_id formData string true "Book ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /books/borrow [post]
func HandleBorrowBook(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	authConn, err := grpc.NewClient("auth-service:50054", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer authConn.Close()

	authClient := authpb.NewAuthServiceClient(authConn)

	userID, _, err := authenticateUser(c.Request.Context(), authClient, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	bookID := c.PostForm("book_id")

	conn, err := grpc.NewClient("borrow-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := borrowpb.NewBorrowServiceClient(conn)

	err = borrowBook(c.Request.Context(), client, userID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to borrow book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book borrowed successfully"})
}

func borrowBook(ctx context.Context, client borrowpb.BorrowServiceClient, userID, bookID string) error {
	req := &borrowpb.BorrowBookRequest{
		UserId: userID,
		BookId: bookID,
	}

	_, err := client.BorrowBook(ctx, req)
	return err
}

// @Summary Return a book
// @Description Return a borrowed book to the library
// @Tags borrow
// @Accept json
// @Produce json
// @Param book_id formData string true "Book ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /books/return [post]
func HandleReturnBook(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	authConn, err := grpc.NewClient("auth-service:50054", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer authConn.Close()

	authClient := authpb.NewAuthServiceClient(authConn)

	userID, _, err := authenticateUser(c.Request.Context(), authClient, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	bookID := c.PostForm("book_id")

	conn, err := grpc.NewClient("borrow-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := borrowpb.NewBorrowServiceClient(conn)

	err = returnBook(c.Request.Context(), client, userID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to return book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}

func returnBook(ctx context.Context, client borrowpb.BorrowServiceClient, userID, bookID string) error {
	req := &borrowpb.ReturnBookRequest{
		UserId: userID,
		BookId: bookID,
	}

	_, err := client.ReturnBook(ctx, req)
	return err
}
