package application

import (
	"context"

	"github.com/hossein-225/Library-Management/book-service/internal/domain"
	pb "github.com/hossein-225/Library-Management/book-service/proto"
)

type BookService struct {
	repo domain.BookRepository
}

func NewBookService(repo domain.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) ListBooks(ctx context.Context) ([]*domain.Book, error) {
	return s.repo.ListBooks(ctx)
}

func (s *BookService) AddBook(ctx context.Context, book *domain.Book) error {
	return s.repo.AddBook(ctx, book)
}

func (s *BookService) UpdateBook(ctx context.Context, book *domain.Book) error {
	return s.repo.UpdateBook(ctx, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) error {
	return s.repo.DeleteBook(ctx, id)
}

func (s *BookService) SearchBooks(ctx context.Context, title, author, category string) ([]*domain.Book, error) {
	return s.repo.SearchBooks(ctx, title, author, category)
}

func (s *BookService) CheckAvailability(ctx context.Context, bookID string) (pb.BookStatus, error) {
	return s.repo.CheckAvailability(ctx, bookID)
}

func (s *BookService) UpdateBookStatus(ctx context.Context, bookID string, status pb.BookStatus) error {
	return s.repo.UpdateBookStatus(ctx, bookID, status)
}
