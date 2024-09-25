package grpc

import (
	"context"

	"github.com/hossein-225/Library-Management/borrow-service/internal/application"
	pb "github.com/hossein-225/Library-Management/borrow-service/proto"
)

type BorrowGRPCServer struct {
	pb.UnimplementedBorrowServiceServer
	service *application.BorrowService
}

func NewBorrowGRPCServer(service *application.BorrowService) *BorrowGRPCServer {
	return &BorrowGRPCServer{service: service}
}

func (s *BorrowGRPCServer) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	err := s.service.BorrowBook(ctx, req.UserId, req.BookId)
	if err != nil {
		return nil, err
	}
	return &pb.BorrowBookResponse{Message: "Book borrowed successfully"}, nil
}

func (s *BorrowGRPCServer) ReturnBook(ctx context.Context, req *pb.ReturnBookRequest) (*pb.ReturnBookResponse, error) {
	err := s.service.ReturnBook(ctx, req.UserId, req.BookId)
	if err != nil {
		return nil, err
	}
	return &pb.ReturnBookResponse{Message: "Book returned successfully"}, nil
}

func (s *BorrowGRPCServer) GetUserBorrows(ctx context.Context, req *pb.GetUserBorrowsRequest) (*pb.GetUserBorrowsResponse, error) {
	borrows, err := s.service.GetUserBorrows(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var response pb.GetUserBorrowsResponse
	for _, borrow := range borrows {
		response.Borrows = append(response.Borrows, &pb.Borrow{
			Id:       borrow.ID,
			UserId:   borrow.UserID,
			BookId:   borrow.BookID,
			Borrowed: borrow.Borrowed,
		})
	}

	return &response, nil
}
