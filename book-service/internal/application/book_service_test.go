package application_test

import (
	"context"
	"testing"

	"github.com/hossein-225/Library-Management/book-service/internal/application"
	"github.com/hossein-225/Library-Management/book-service/internal/domain"
	pb "github.com/hossein-225/Library-Management/book-service/proto"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	repo := &MockBookRepository{}
	service := application.NewBookService(repo)

	book := &domain.Book{
		ID:       "123",
		Title:    "Test Book",
		Author:   "John Doe",
		Category: "Fiction",
	}

	err := service.AddBook(context.Background(), book)

	assert.NoError(t, err)
	assert.Equal(t, book, repo.AddedBook)
}

func TestListBooks(t *testing.T) {
	repo := &MockBookRepository{
		Books: []*domain.Book{
			{
				ID:       "123",
				Title:    "Test Book",
				Author:   "John Doe",
				Category: "Fiction",
				Status:   pb.BookStatus_AVAILABLE,
			},
		},
	}
	service := application.NewBookService(repo)

	books, err := service.ListBooks(context.Background())

	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.Equal(t, "Test Book", books[0].Title)
}

type MockBookRepository struct {
	Books     []*domain.Book
	AddedBook *domain.Book
}

func (m *MockBookRepository) CheckAvailability(ctx context.Context, bookID string) (pb.BookStatus, error) {
	panic("unimplemented")
}

func (m *MockBookRepository) UpdateBookStatus(ctx context.Context, bookID string, status pb.BookStatus) error {
	panic("unimplemented")
}

func (m *MockBookRepository) SearchBooks(ctx context.Context, title string, author string, category string) ([]*domain.Book, error) {
	panic("unimplemented")
}

func (m *MockBookRepository) ListBooks(ctx context.Context) ([]*domain.Book, error) {
	return m.Books, nil
}

func (m *MockBookRepository) AddBook(ctx context.Context, book *domain.Book) error {
	m.AddedBook = book
	return nil
}

func (m *MockBookRepository) UpdateBook(ctx context.Context, book *domain.Book) error {
	return nil
}

func (m *MockBookRepository) DeleteBook(ctx context.Context, id string) error {
	return nil
}
