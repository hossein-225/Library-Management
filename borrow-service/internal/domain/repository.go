package domain

type BorrowRepository interface {
	BorrowBook(borrow *Borrow) error
	ReturnBook(userID, bookID string) error
	GetUserBorrows(userID string) ([]*Borrow, error)
}
