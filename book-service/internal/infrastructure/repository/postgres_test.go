package repository_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hossein-225/Library-Management/book-service/internal/domain"
	"github.com/hossein-225/Library-Management/book-service/internal/infrastructure/repository"
	pb "github.com/hossein-225/Library-Management/book-service/proto"
	"github.com/stretchr/testify/assert"
)

func TestPostgresBookRepository_ListBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewPostgresBookRepository(db)

	rows := sqlmock.NewRows([]string{"id", "title", "author", "category", "status"}).
		AddRow("1", "Book 1", "Author 1", "Category 1", pb.BookStatus_AVAILABLE).
		AddRow("2", "Book 2", "Author 2", "Category 2", pb.BookStatus_BORROWED)

	mock.ExpectQuery("SELECT id, title, author, category, status FROM books").
		WillReturnRows(rows)

	books, err := repo.ListBooks(context.Background())
	assert.NoError(t, err)
	assert.Len(t, books, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresBookRepository_AddBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewPostgresBookRepository(db)

	book := &domain.Book{
		ID:       "1",
		Title:    "Book 1",
		Author:   "Author 1",
		Category: "Category 1",
		Status:   pb.BookStatus_AVAILABLE,
	}

	mock.ExpectExec("INSERT INTO books").
		WithArgs(book.ID, book.Title, book.Author, book.Category, pb.BookStatus_AVAILABLE).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddBook(context.Background(), book)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresBookRepository_UpdateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewPostgresBookRepository(db)

	book := &domain.Book{
		ID:       "1",
		Title:    "Updated Book",
		Author:   "Updated Author",
		Category: "Updated Category",
	}

	mock.ExpectExec("UPDATE books SET title").
		WithArgs(book.Title, book.Author, book.Category, book.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateBook(context.Background(), book)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresBookRepository_DeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewPostgresBookRepository(db)

	mock.ExpectExec("DELETE FROM books WHERE id").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteBook(context.Background(), "1")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresBookRepository_SearchBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewPostgresBookRepository(db)

	rows := sqlmock.NewRows([]string{"id", "title", "author", "category", "status"}).
		AddRow("1", "Book 1", "Author 1", "Category 1", pb.BookStatus_AVAILABLE)

	mock.ExpectQuery("SELECT id, title, author, category, status FROM books WHERE").
		WithArgs("%Book%", "%Author%", "%Category%").
		WillReturnRows(rows)

	books, err := repo.SearchBooks(context.Background(), "Book", "Author", "Category")
	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresBookRepository_CheckAvailability(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewPostgresBookRepository(db)

	mock.ExpectQuery("SELECT status FROM books WHERE id").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"status"}).AddRow(pb.BookStatus_AVAILABLE))

	status, err := repo.CheckAvailability(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, pb.BookStatus_AVAILABLE, status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresBookRepository_UpdateBookStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewPostgresBookRepository(db)

	mock.ExpectExec("UPDATE books SET status").
		WithArgs(pb.BookStatus_BORROWED, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateBookStatus(context.Background(), "1", pb.BookStatus_BORROWED)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
