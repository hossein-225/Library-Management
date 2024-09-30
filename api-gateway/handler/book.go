package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	authpb "github.com/hossein-225/Library-Management/auth-service/proto"
	bookpb "github.com/hossein-225/Library-Management/book-service/proto"
	"google.golang.org/grpc"
)

// @Summary List books
// @Description Get a list of all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /books [get]
func HandleBookList(c *gin.Context) {
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

	_, _, err = authenticateUser(c.Request.Context(), authClient, token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	conn, err := grpc.NewClient("book-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)

	books, err := fetchBooks(c.Request.Context(), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func fetchBooks(ctx context.Context, client bookpb.BookServiceClient) ([]*bookpb.Book, error) {
	req := &bookpb.ListBooksRequest{}
	res, err := client.ListBooks(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Books, nil
}

// @Summary Add a book
// @Description Add a new book (Admins only)
// @Tags books
// @Accept json
// @Produce json
// @Param title formData string true "Book title"
// @Param author formData string true "Book author"
// @Param category formData string true "Book category"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /books [post]
func HandleAddBook(c *gin.Context) {
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

	_, _, err = authenticateUser(c.Request.Context(), authClient, token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	title := c.PostForm("title")
	author := c.PostForm("author")
	category := c.PostForm("category")

	conn, err := grpc.NewClient("book-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)

	err = addBook(c.Request.Context(), client, title, author, category)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})
}

func addBook(ctx context.Context, client bookpb.BookServiceClient, title, author, category string) error {
	req := &bookpb.AddBookRequest{
		Title:    title,
		Author:   author,
		Category: category,
	}

	_, err := client.AddBook(ctx, req)
	return err
}

// @Summary Update a book
// @Description Update a book's information (Admins only)
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param title formData string false "Book title"
// @Param author formData string false "Book author"
// @Param category formData string false "Book category"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /books/{id} [put]
func HandleUpdateBook(c *gin.Context) {
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

	_, isAdmin, err := authenticateUser(c.Request.Context(), authClient, token)
	if err != nil || !isAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Admins only"})
		return
	}

	bookID := c.Param("id")

	title := c.PostForm("title")
	author := c.PostForm("author")
	category := c.PostForm("category")

	conn, err := grpc.NewClient("book-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)

	err = updateBook(c.Request.Context(), client, bookID, title, author, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

func updateBook(ctx context.Context, client bookpb.BookServiceClient, bookID string, title, author, category string) error {

	req := &bookpb.UpdateBookRequest{
		Id:       bookID,
		Title:    title,
		Author:   author,
		Category: category,
	}

	_, err := client.UpdateBook(ctx, req)
	return err
}

// @Summary Delete a book
// @Description Delete a book from the library (Admins only)
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /books/{id} [delete]
func HandleDeleteBook(c *gin.Context) {
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

	_, isAdmin, err := authenticateUser(c.Request.Context(), authClient, token)
	if err != nil || !isAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Admins only"})
		return
	}

	bookID := c.Param("id")

	conn, err := grpc.NewClient("book-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)

	err = deleteBook(c.Request.Context(), client, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func deleteBook(ctx context.Context, client bookpb.BookServiceClient, bookID string) error {
	req := &bookpb.DeleteBookRequest{Id: bookID}

	_, err := client.DeleteBook(ctx, req)
	return err
}

// @Summary Search books
// @Description Search books by title, author, or category
// @Tags books
// @Accept json
// @Produce json
// @Param title query string false "Book title"
// @Param author query string false "Book author"
// @Param category query string false "Book category"
// @Success 200 {array} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /books/search [get]
func HandleSearchBooks(c *gin.Context) {
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

	_, _, err = authenticateUser(c.Request.Context(), authClient, token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	title := c.Query("title")
	author := c.Query("author")
	category := c.Query("category")

	conn, err := grpc.NewClient("book-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)

	books, err := searchBooks(c.Request.Context(), client, title, author, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func searchBooks(ctx context.Context, client bookpb.BookServiceClient, title, author, category string) ([]*bookpb.Book, error) {
	req := &bookpb.SearchBooksRequest{
		Title:    title,
		Author:   author,
		Category: category,
	}

	res, err := client.SearchBooks(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Books, nil
}
