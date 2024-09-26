package domain

type BookRepository interface {
	ListBooks() ([]*Book, error)
	AddBook(book *Book) error
	UpdateBook(book *Book) error
	DeleteBook(id string) error
	SearchBooks(title, author, category string) ([]*Book, error)
}
