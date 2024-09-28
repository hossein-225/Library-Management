package grpc

import (
	"context"
	"log"

	"github.com/hossein-225/Library-Management/book-service/pkg/utils"
	pb "github.com/hossein-225/Library-Management/book-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hossein-225/Library-Management/book-service/internal/application"
	"github.com/hossein-225/Library-Management/book-service/internal/domain"
)

type BookGRPCServer struct {
	pb.UnimplementedBookServiceServer
	service *application.BookService
}

func NewBookGRPCServer(service *application.BookService) *BookGRPCServer {
	return &BookGRPCServer{service: service}
}

// ListBooks
// @Summary Retrieve a list of books
// @Description Retrieves a list of available books in the system
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "List of books retrieved successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /books [get]
func (s *BookGRPCServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	books, err := s.service.ListBooks(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve list of books: %v", err)
	}

	var response pb.ListBooksResponse
	for _, book := range books {
		response.Books = append(response.Books, &pb.Book{
			Id:       book.ID,
			Title:    book.Title,
			Author:   book.Author,
			Category: book.Category,
			Status:   book.Status,
		})
	}

	return &response, nil
}

// AddBook
// @Summary Add a new book
// @Description Adds a new book to the library
// @Tags books
// @Accept  json
// @Produce  json
// @Param   title     body   string   true  "Title of the book"
// @Param   author    body   string   true  "Author of the book"
// @Param   category  body   string   true  "Category of the book"
// @Success 200 {object} map[string]interface{} "Book added successfully"
// @Failure 400 {string} string "Title, author, or category cannot be empty"
// @Failure 500 {string} string "Internal server error"
// @Router /books [post]
func (s *BookGRPCServer) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	if req.Title == "" || req.Author == "" || req.Category == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Title, author, or category cannot be empty")
	}

	book := &domain.Book{
		ID:       utils.GenerateUUID(),
		Title:    req.Title,
		Author:   req.Author,
		Category: req.Category,
		Status:   pb.BookStatus_AVAILABLE,
	}

	if err := s.service.AddBook(ctx, book); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add book: %v", err)
	}

	return &pb.AddBookResponse{
		Book: &pb.Book{
			Id:       book.ID,
			Title:    book.Title,
			Author:   book.Author,
			Category: book.Category,
			Status:   book.Status,
		},
	}, nil
}

// UpdateBook
// @Summary Update a book's information
// @Description Updates the information of an existing book in the library
// @Tags books
// @Accept  json
// @Produce  json
// @Param   id        body   string   true  "ID of the book"
// @Param   title     body   string   true  "Title of the book"
// @Param   author    body   string   true  "Author of the book"
// @Param   category  body   string   true  "Category of the book"
// @Param   available body   bool     true  "Availability status of the book"
// @Success 200 {object} map[string]interface{} "Book updated successfully"
// @Failure 400 {string} string "ID, title, author, or category cannot be empty"
// @Failure 404 {string} string "Book not found"
// @Failure 500 {string} string "Internal server error"
// @Router /books/{id} [put]
func (s *BookGRPCServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	if req.Id == "" || req.Title == "" || req.Author == "" || req.Category == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ID, title, author, or category cannot be empty")
	}

	book := &domain.Book{
		ID:       req.Id,
		Title:    req.Title,
		Author:   req.Author,
		Category: req.Category,
	}

	if err := s.service.UpdateBook(ctx, book); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update book: %v", err)
	}

	return &pb.UpdateBookResponse{
		Book: &pb.Book{
			Id:       book.ID,
			Title:    book.Title,
			Author:   book.Author,
			Category: book.Category,
		},
	}, nil
}

// DeleteBook
// @Summary Delete a book
// @Description Deletes a book from the library by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param   id  path   string   true  "ID of the book to delete"
// @Success 200 {object} map[string]interface{} "Book deleted successfully"
// @Failure 400 {string} string "Book ID cannot be empty"
// @Failure 404 {string} string "Book not found"
// @Failure 500 {string} string "Internal server error"
// @Router /books/{id} [delete]
func (s *BookGRPCServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Book ID cannot be empty")
	}

	if err := s.service.DeleteBook(ctx, req.Id); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to delete book: %v", err)
	}

	return &pb.DeleteBookResponse{
		Message: "Book deleted successfully",
	}, nil
}

func (s *BookGRPCServer) SearchBooks(ctx context.Context, req *pb.SearchBooksRequest) (*pb.SearchBooksResponse, error) {
	books, err := s.service.SearchBooks(ctx, req.Title, req.Author, req.Category)
	if err != nil {
		return nil, err
	}

	var response pb.SearchBooksResponse
	for _, book := range books {
		response.Books = append(response.Books, &pb.Book{
			Id:       book.ID,
			Title:    book.Title,
			Author:   book.Author,
			Category: book.Category,
			Status:   book.Status,
		})
	}

	return &response, nil
}

// func (s *BookGRPCServer) CheckAvailability(ctx context.Context, req *pb.CheckAvailabilityRequest) (*pb.CheckAvailabilityResponse, error) {
// 	book, err := s.service.GetBookByID(req.BookId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.CheckAvailabilityResponse{Status: book.Available}, nil
// }

// func (s *BookGRPCServer) UpdateBookStatus(ctx context.Context, req *pb.UpdateBookStatusRequest) (*pb.UpdateBookStatusResponse, error) {
// 	err := s.service.UpdateBookStatus(req.BookId, req.Status)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.UpdateBookStatusResponse{Success: true}, nil
// }
