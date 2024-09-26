package grpc

import (
	"context"

	"github.com/hossein-225/Library-Management/borrow-service/internal/application"
	pb "github.com/hossein-225/Library-Management/borrow-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BorrowGRPCServer struct {
	pb.UnimplementedBorrowServiceServer
	service *application.BorrowService
}

func NewBorrowGRPCServer(service *application.BorrowService) *BorrowGRPCServer {
	return &BorrowGRPCServer{service: service}
}

// BorrowBook
// @Summary Borrow a book
// @Description Allows a user to borrow a book from the library
// @Tags borrow
// @Accept  json
// @Produce  json
// @Param   user_id  body   string   true  "User ID"
// @Param   book_id  body   string   true  "Book ID"
// @Success 200 {object} map[string]interface{} "Book borrowed successfully"
// @Failure 400 {string} string "User ID or Book ID cannot be empty"
// @Failure 404 {string} string "Book or User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /borrow [post]
func (s *BorrowGRPCServer) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	if req.UserId == "" || req.BookId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "User ID or Book ID cannot be empty")
	}

	err := s.service.BorrowBook(ctx, req.UserId, req.BookId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to borrow book: %v", err)
	}

	return &pb.BorrowBookResponse{
		Message: "Book borrowed successfully",
	}, nil
}

// ReturnBook
// @Summary Return a borrowed book
// @Description Allows a user to return a borrowed book to the library
// @Tags borrow
// @Accept  json
// @Produce  json
// @Param   user_id  body   string   true  "User ID"
// @Param   book_id  body   string   true  "Book ID"
// @Success 200 {object} map[string]interface{} "Book returned successfully"
// @Failure 400 {string} string "User ID or Book ID cannot be empty"
// @Failure 404 {string} string "Book or User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /return [post]
func (s *BorrowGRPCServer) ReturnBook(ctx context.Context, req *pb.ReturnBookRequest) (*pb.ReturnBookResponse, error) {
	if req.UserId == "" || req.BookId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "User ID or Book ID cannot be empty")
	}

	err := s.service.ReturnBook(ctx, req.UserId, req.BookId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to return book: %v", err)
	}

	return &pb.ReturnBookResponse{
		Message: "Book returned successfully",
	}, nil
}

// GetUserBorrows
// @Summary Get user's borrowed books
// @Description Retrieves the list of books borrowed by a specific user
// @Tags borrow
// @Accept  json
// @Produce  json
// @Param   user_id  body   string   true  "User ID"
// @Success 200 {object} map[string]interface{} "List of borrowed books retrieved successfully"
// @Failure 400 {string} string "User ID cannot be empty"
// @Failure 404 {string} string "No borrow records found for this user"
// @Failure 500 {string} string "Internal server error"
// @Router /borrows/user [get]
func (s *BorrowGRPCServer) GetUserBorrows(ctx context.Context, req *pb.GetUserBorrowsRequest) (*pb.GetUserBorrowsResponse, error) {
	if req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "User ID cannot be empty")
	}

	borrows, err := s.service.GetUserBorrows(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve borrow records: %v", err)
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
