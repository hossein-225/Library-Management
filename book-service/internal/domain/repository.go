package domain

import (
	"context"

	pb "github.com/hossein-225/Library-Management/book-service/proto"
)

type BookRepository interface {
	ListBooks(ctx context.Context) ([]*Book, error)
	AddBook(ctx context.Context, book *Book) error
	UpdateBook(ctx context.Context, book *Book) error
	DeleteBook(ctx context.Context, id string) error
	SearchBooks(ctx context.Context, title, author, category string) ([]*Book, error)
	CheckAvailability(ctx context.Context, bookID string) (pb.BookStatus, error)
	UpdateBookStatus(ctx context.Context, bookID string, status pb.BookStatus) error
}
