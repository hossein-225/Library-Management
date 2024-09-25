package repository

import (
	"database/sql"

	"github.com/hossein-225/Library-Management/borrow-service/internal/domain"
)

type PostgresBorrowRepository struct {
	db *sql.DB
}

func NewPostgresBorrowRepository(db *sql.DB) *PostgresBorrowRepository {
	return &PostgresBorrowRepository{db: db}
}

func (r *PostgresBorrowRepository) BorrowBook(borrow *domain.Borrow) error {
	_, err := r.db.Exec("INSERT INTO borrows (id, user_id, book_id, borrowed) VALUES ($1, $2, $3, $4)",
		borrow.ID, borrow.UserID, borrow.BookID, borrow.Borrowed)
	return err
}

func (r *PostgresBorrowRepository) ReturnBook(userID, bookID string) error {
	_, err := r.db.Exec("UPDATE borrows SET borrowed = false WHERE user_id = $1 AND book_id = $2", userID, bookID)
	return err
}

func (r *PostgresBorrowRepository) GetUserBorrows(userID string) ([]*domain.Borrow, error) {
	rows, err := r.db.Query("SELECT id, user_id, book_id, borrowed FROM borrows WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var borrows []*domain.Borrow
	for rows.Next() {
		borrow := &domain.Borrow{}
		if err := rows.Scan(&borrow.ID, &borrow.UserID, &borrow.BookID, &borrow.Borrowed); err != nil {
			return nil, err
		}
		borrows = append(borrows, borrow)
	}
	return borrows, nil
}
