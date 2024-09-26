package repository

import (
	"database/sql"

	"github.com/hossein-225/Library-Management/book-service/internal/domain"
)

type PostgresBookRepository struct {
	db *sql.DB
}

func NewPostgresBookRepository(db *sql.DB) *PostgresBookRepository {
	return &PostgresBookRepository{db: db}
}

func (r *PostgresBookRepository) ListBooks() ([]*domain.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, category, available FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*domain.Book
	for rows.Next() {
		book := &domain.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Available); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *PostgresBookRepository) AddBook(book *domain.Book) error {
	_, err := r.db.Exec("INSERT INTO books (id, title, author, category, available) VALUES ($1, $2, $3, $4, $5)", book.ID, book.Title, book.Author, book.Category, book.Available)
	return err
}

func (r *PostgresBookRepository) UpdateBook(book *domain.Book) error {
	_, err := r.db.Exec("UPDATE books SET title=$1, author=$2, category=$3, available=$4 WHERE id=$5", book.Title, book.Author, book.Category, book.Available, book.ID)
	return err
}

func (r *PostgresBookRepository) DeleteBook(id string) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}

func (r *PostgresBookRepository) SearchBooks(title, author, category string) ([]*domain.Book, error) {
	query := "SELECT id, title, author, category, available FROM books WHERE 1=1"
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
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Available); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
