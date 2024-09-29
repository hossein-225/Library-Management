package domain

import "time"

type Borrow struct {
	ID         string
	UserID     string
	BookID     string
	Borrowed   bool
	BorrowedAt time.Time
	ReturnedAt *time.Time
}
