package grpc

import (
	"context"

	"github.com/hossein-225/Library-Management/book-service/pkg/utils"
	pb "github.com/hossein-225/Library-Management/book-service/proto"

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

func (s *BookGRPCServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	books, err := s.service.ListBooks(ctx)
	if err != nil {
		return nil, err
	}

	var response pb.ListBooksResponse
	for _, book := range books {
		response.Books = append(response.Books, &pb.Book{
			Id:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Category:  book.Category,
			Available: book.Available,
		})
	}

	return &response, nil
}

func (s *BookGRPCServer) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	book := &domain.Book{
		ID:        utils.GenerateUUID(),
		Title:     req.Title,
		Author:    req.Author,
		Category:  req.Category,
		Available: true,
	}

	if err := s.service.AddBook(ctx, book); err != nil {
		return nil, err
	}

	return &pb.AddBookResponse{
		Book: &pb.Book{
			Id:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Category:  book.Category,
			Available: book.Available,
		},
	}, nil
}

func (s *BookGRPCServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	book := &domain.Book{
		ID:        req.Id,
		Title:     req.Title,
		Author:    req.Author,
		Category:  req.Category,
		Available: req.Available,
	}

	if err := s.service.UpdateBook(ctx, book); err != nil {
		return nil, err
	}

	return &pb.UpdateBookResponse{
		Book: &pb.Book{
			Id:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Category:  book.Category,
			Available: book.Available,
		},
	}, nil
}

func (s *BookGRPCServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	if err := s.service.DeleteBook(ctx, req.Id); err != nil {
		return nil, err
	}

	return &pb.DeleteBookResponse{
		Message: "Book deleted successfully",
	}, nil
}
