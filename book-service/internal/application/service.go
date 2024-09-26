package application

import (
	"context"

	"github.com/hossein-225/Library-Management/book-service/internal/domain"
)

type BookService struct {
	repo domain.BookRepository
}

func NewBookService(repo domain.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) ListBooks(ctx context.Context) ([]*domain.Book, error) {
	return s.repo.ListBooks()
}

func (s *BookService) AddBook(ctx context.Context, book *domain.Book) error {
	return s.repo.AddBook(book)
}

func (s *BookService) UpdateBook(ctx context.Context, book *domain.Book) error {
	return s.repo.UpdateBook(book)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) error {
	return s.repo.DeleteBook(id)
}

func (s *BookService) SearchBooks(ctx context.Context, title, author, category string) ([]*domain.Book, error) {
	return s.repo.SearchBooks(title, author, category)
}
