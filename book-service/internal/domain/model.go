package domain

import pb "github.com/hossein-225/Library-Management/book-service/proto"

type Book struct {
	ID       string
	Title    string
	Author   string
	Category string
	Status   pb.BookStatus
}
