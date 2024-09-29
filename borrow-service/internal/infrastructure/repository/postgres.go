package repository

import (
	"database/sql"
	"time"

	"github.com/hossein-225/Library-Management/borrow-service/internal/domain"
)

type PostgresBorrowRepository struct {
	db *sql.DB
}

func NewPostgresBorrowRepository(db *sql.DB) *PostgresBorrowRepository {
	return &PostgresBorrowRepository{db: db}
}

func (r *PostgresBorrowRepository) BorrowBook(borrow *domain.Borrow) error {
	borrow.BorrowedAt = time.Now()

	_, err := r.db.Exec("INSERT INTO borrows (id, user_id, book_id, borrowed, borrowed_at) VALUES ($1, $2, $3, $4, $5)",
		borrow.ID, borrow.UserID, borrow.BookID, borrow.Borrowed, borrow.BorrowedAt)
	return err
}

func (r *PostgresBorrowRepository) ReturnBook(userID, bookID string) error {
	returnedAt := time.Now()

	_, err := r.db.Exec("UPDATE borrows SET borrowed = false, returned_at = $1 WHERE user_id = $2 AND book_id = $3",
		returnedAt, userID, bookID)
	return err
}
