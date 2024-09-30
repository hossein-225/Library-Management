package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hossein-225/Library-Management/borrow-service/internal/domain"
	"github.com/hossein-225/Library-Management/borrow-service/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestPostgresBorrowRepository_BorrowBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	borrowRepo := repository.NewPostgresBorrowRepository(db)
	borrow := &domain.Borrow{
		ID:       "borrow123",
		UserID:   "user123",
		BookID:   "book123",
		Borrowed: true,
	}

	mock.ExpectExec("INSERT INTO borrows").
		WithArgs(borrow.ID, borrow.UserID, borrow.BookID, borrow.Borrowed, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = borrowRepo.BorrowBook(borrow)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresBorrowRepository_ReturnBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	borrowRepo := repository.NewPostgresBorrowRepository(db)
	userID := "user123"
	bookID := "book123"

	mock.ExpectExec("UPDATE borrows SET borrowed = false, returned_at =").
		WithArgs(sqlmock.AnyArg(), userID, bookID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = borrowRepo.ReturnBook(userID, bookID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
