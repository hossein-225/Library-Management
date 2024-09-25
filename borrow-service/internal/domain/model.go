package domain

type Borrow struct {
	ID       string
	UserID   string
	BookID   string
	Borrowed bool
}
