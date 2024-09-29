package application_test

import (
	"context"
	"testing"

	"github.com/hossein-225/Library-Management/borrow-service/internal/application"
	"github.com/hossein-225/Library-Management/borrow-service/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestBorrowBook(t *testing.T) {
	repo := &MockBorrowRepository{}
	service := application.NewBorrowService(repo)

	err := service.BorrowBook(context.Background(), "user123", "book456")

	assert.NoError(t, err)
	assert.Equal(t, "user123", repo.BorrowedBook.UserID)
	assert.Equal(t, "book456", repo.BorrowedBook.BookID)
	assert.True(t, repo.BorrowedBook.Borrowed)
}

func TestReturnBook(t *testing.T) {
	repo := &MockBorrowRepository{}
	service := application.NewBorrowService(repo)

	err := service.ReturnBook(context.Background(), "user123", "book456")

	assert.NoError(t, err)
	assert.Equal(t, "user123", repo.ReturnedUserID)
	assert.Equal(t, "book456", repo.ReturnedBookID)
}

type MockBorrowRepository struct {
	Borrows        []*domain.Borrow
	BorrowedBook   *domain.Borrow
	ReturnedUserID string
	ReturnedBookID string
}

func (m *MockBorrowRepository) BorrowBook(borrow *domain.Borrow) error {
	m.BorrowedBook = borrow
	return nil
}

func (m *MockBorrowRepository) ReturnBook(userID, bookID string) error {
	m.ReturnedUserID = userID
	m.ReturnedBookID = bookID
	return nil
}

func (m *MockBorrowRepository) GetUserBorrows(userID string) ([]*domain.Borrow, error) {
	return m.Borrows, nil
}
