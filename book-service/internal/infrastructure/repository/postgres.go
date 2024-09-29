package repository

import (
	"context"
	"database/sql"

	"github.com/hossein-225/Library-Management/book-service/internal/domain"
	pb "github.com/hossein-225/Library-Management/book-service/proto"
)

type PostgresBookRepository struct {
	db *sql.DB
}

func NewPostgresBookRepository(db *sql.DB) *PostgresBookRepository {
	return &PostgresBookRepository{db: db}
}

func (r *PostgresBookRepository) ListBooks(ctx context.Context) ([]*domain.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, category, status FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*domain.Book
	for rows.Next() {
		book := &domain.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Status); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *PostgresBookRepository) AddBook(ctx context.Context, book *domain.Book) error {
	_, err := r.db.Exec("INSERT INTO books (id, title, author, category, status) VALUES ($1, $2, $3, $4, $5)", book.ID, book.Title, book.Author, book.Category, pb.BookStatus_AVAILABLE)

	return err
}

func (r *PostgresBookRepository) UpdateBook(ctx context.Context, book *domain.Book) error {
	_, err := r.db.Exec("UPDATE books SET title=$1, author=$2, category=$3 WHERE id=$4", book.Title, book.Author, book.Category, book.ID)

	return err
}

func (r *PostgresBookRepository) DeleteBook(ctx context.Context, id string) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id=$1", id)

	return err
}

func (r *PostgresBookRepository) SearchBooks(ctx context.Context, title, author, category string) ([]*domain.Book, error) {
	query := "SELECT id, title, author, category, status FROM books WHERE 1=1"
	params := []interface{}{}

	if title != "" {
		query += " AND title ILIKE $1"
		params = append(params, "%"+title+"%")
	}
	if author != "" {
		query += " AND author ILIKE $2"
		params = append(params, "%"+author+"%")
	}
	if category != "" {
		query += " AND category ILIKE $3"
		params = append(params, "%"+category+"%")
	}

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*domain.Book
	for rows.Next() {
		book := &domain.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Status); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *PostgresBookRepository) CheckAvailability(ctx context.Context, bookID string) (pb.BookStatus, error) {
	var status pb.BookStatus
	err := r.db.QueryRow("SELECT status FROM books WHERE id=$1", bookID).Scan(&status)
	if err != nil {
		return pb.BookStatus_BOOK_STATUS_UNSPECIFIED, err
	}

	return status, nil
}

func (r *PostgresBookRepository) UpdateBookStatus(ctx context.Context, bookID string, status pb.BookStatus) error {
	_, err := r.db.Exec("UPDATE books SET status=$1 WHERE id=$2", status, bookID)

	return err
}
