package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/hossein-225/Library-Management/borrow-service/internal/application"
	"github.com/hossein-225/Library-Management/borrow-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBorrowRepository struct {
	mock.Mock
}

func (m *MockBorrowRepository) BorrowBook(borrow *domain.Borrow) error {
	args := m.Called(borrow)
	return args.Error(0)
}

func (m *MockBorrowRepository) ReturnBook(userID, bookID string) error {
	args := m.Called(userID, bookID)
	return args.Error(0)
}

func (m *MockBorrowRepository) GetUserBorrows(userID string) ([]*domain.Borrow, error) {
	args := m.Called(userID)
	return args.Get(0).([]*domain.Borrow), args.Error(1)
}

func TestBorrowBook_Success(t *testing.T) {
	repo := new(MockBorrowRepository)
	service := application.NewBorrowService(repo)

	repo.On("BorrowBook", mock.MatchedBy(func(b *domain.Borrow) bool {
		return b.UserID == "user123" && b.BookID == "book456" && b.Borrowed == true
	})).Return(nil)

	err := service.BorrowBook(context.Background(), "user123", "book456")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestBorrowBook_Failure(t *testing.T) {
	repo := new(MockBorrowRepository)
	service := application.NewBorrowService(repo)

	repo.On("BorrowBook", mock.Anything).Return(errors.New("borrow failed"))

	err := service.BorrowBook(context.Background(), "user123", "book456")

	assert.Error(t, err)
	assert.Equal(t, "borrow failed", err.Error())
	repo.AssertExpectations(t)
}

func TestReturnBook_Success(t *testing.T) {
	repo := new(MockBorrowRepository)
	service := application.NewBorrowService(repo)

	repo.On("ReturnBook", "user123", "book456").Return(nil)

	err := service.ReturnBook(context.Background(), "user123", "book456")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestReturnBook_Failure(t *testing.T) {
	repo := new(MockBorrowRepository)
	service := application.NewBorrowService(repo)

	repo.On("ReturnBook", "user123", "book456").Return(errors.New("return failed"))

	err := service.ReturnBook(context.Background(), "user123", "book456")

	assert.Error(t, err)
	assert.Equal(t, "return failed", err.Error())
	repo.AssertExpectations(t)
}
