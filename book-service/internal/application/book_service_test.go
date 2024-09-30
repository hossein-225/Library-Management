package application_test

import (
	"context"
	"testing"

	"github.com/hossein-225/Library-Management/book-service/internal/application"
	"github.com/hossein-225/Library-Management/book-service/internal/domain"
	pb "github.com/hossein-225/Library-Management/book-service/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestUpdateBook(t *testing.T) {
	repo := &MockBookRepository{}
	service := application.NewBookService(repo)

	book := &domain.Book{
		ID:       "123",
		Title:    "Updated Book",
		Author:   "Jane Doe",
		Category: "Science",
	}

	repo.On("UpdateBook", book).Return(nil)

	err := service.UpdateBook(context.Background(), book)

	assert.NoError(t, err)
	repo.AssertCalled(t, "UpdateBook", book)
}

func TestDeleteBook(t *testing.T) {
	repo := &MockBookRepository{}
	service := application.NewBookService(repo)

	repo.On("DeleteBook", "123").Return(nil)

	err := service.DeleteBook(context.Background(), "123")

	assert.NoError(t, err)
	repo.AssertCalled(t, "DeleteBook", "123")
}

func TestSearchBooks(t *testing.T) {
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

	repo.On("SearchBooks", mock.Anything, "Test Book", "John Doe", "Fiction").Return(repo.Books, nil)

	books, err := service.SearchBooks(context.Background(), "Test Book", "John Doe", "Fiction")

	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.Equal(t, "Test Book", books[0].Title)
	assert.Equal(t, "John Doe", books[0].Author)
}

func TestCheckAvailability(t *testing.T) {
	repo := &MockBookRepository{}
	service := application.NewBookService(repo)

	repo.On("CheckAvailability", mock.Anything, "123").Return(pb.BookStatus_AVAILABLE, nil)

	status, err := service.CheckAvailability(context.Background(), "123")

	assert.NoError(t, err)
	assert.Equal(t, pb.BookStatus_AVAILABLE, status)
}

func TestUpdateBookStatus(t *testing.T) {
	repo := &MockBookRepository{}
	service := application.NewBookService(repo)

	repo.On("UpdateBookStatus", "123", pb.BookStatus_BORROWED).Return(nil)

	err := service.UpdateBookStatus(context.Background(), "123", pb.BookStatus_BORROWED)

	assert.NoError(t, err)
	repo.AssertCalled(t, "UpdateBookStatus", "123", pb.BookStatus_BORROWED)
}

type MockBookRepository struct {
	mock.Mock
	Books     []*domain.Book
	AddedBook *domain.Book
}

func (m *MockBookRepository) CheckAvailability(ctx context.Context, bookID string) (pb.BookStatus, error) {
	args := m.Called(ctx, bookID)
	return args.Get(0).(pb.BookStatus), args.Error(1)
}

func (m *MockBookRepository) UpdateBookStatus(ctx context.Context, bookID string, status pb.BookStatus) error {
	args := m.Called(bookID, status)
	return args.Error(0)
}

func (m *MockBookRepository) SearchBooks(ctx context.Context, title, author, category string) ([]*domain.Book, error) {
	args := m.Called(ctx, title, author, category)
	return m.Books, args.Error(1)
}

func (m *MockBookRepository) ListBooks(ctx context.Context) ([]*domain.Book, error) {
	return m.Books, nil
}

func (m *MockBookRepository) AddBook(ctx context.Context, book *domain.Book) error {
	m.AddedBook = book
	return nil
}

func (m *MockBookRepository) UpdateBook(ctx context.Context, book *domain.Book) error {
	m.Called(book)
	return nil
}

func (m *MockBookRepository) DeleteBook(ctx context.Context, id string) error {
	m.Called(id)
	return nil
}
