package application

import (
	"context"

	"github.com/hossein-225/Library-Management/borrow-service/internal/domain"
	"github.com/hossein-225/Library-Management/borrow-service/pkg/utils"
)

type BorrowService struct {
	repo domain.BorrowRepository
}

func NewBorrowService(repo domain.BorrowRepository) *BorrowService {
	return &BorrowService{repo: repo}
}

func (s *BorrowService) BorrowBook(ctx context.Context, userID, bookID string) error {
	borrow := &domain.Borrow{
		ID:       utils.GenerateUUID(),
		UserID:   userID,
		BookID:   bookID,
		Borrowed: true,
	}
	return s.repo.BorrowBook(borrow)
}

func (s *BorrowService) ReturnBook(ctx context.Context, userID, bookID string) error {
	return s.repo.ReturnBook(userID, bookID)
}
