package grpc

import (
	"context"

	bookpb "github.com/hossein-225/Library-Management/book-service/proto"
	"github.com/hossein-225/Library-Management/borrow-service/internal/application"
	pb "github.com/hossein-225/Library-Management/borrow-service/proto"
	"google.golang.org/grpc"
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

	conn, err := grpc.NewClient("book-service:50051", grpc.WithInsecure())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to connect to book service: %v", err)
	}
	defer conn.Close()

	bookClient := bookpb.NewBookServiceClient(conn)

	checkReq := &bookpb.CheckAvailabilityRequest{BookId: req.BookId}
	checkRes, err := bookClient.CheckAvailability(ctx, checkReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to check book availability: %v", err)
	}

	if checkRes.Status != bookpb.BookStatus_AVAILABLE {
		return nil, status.Errorf(codes.FailedPrecondition, "Book is not available for borrowing")
	}

	err = s.service.BorrowBook(ctx, req.UserId, req.BookId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to borrow book: %v", err)
	}

	updateStatusReq := &bookpb.UpdateBookStatusRequest{
		BookId: req.BookId,
		Status: bookpb.BookStatus_BORROWED,
	}
	_, err = bookClient.UpdateBookStatus(ctx, updateStatusReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update book status: %v", err)
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

	conn, err := grpc.NewClient("book-service:50051", grpc.WithInsecure())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to connect to book service: %v", err)
	}
	defer conn.Close()

	bookClient := bookpb.NewBookServiceClient(conn)

	updateStatusReq := &bookpb.UpdateBookStatusRequest{
		BookId: req.BookId,
		Status: bookpb.BookStatus_AVAILABLE,
	}
	_, err = bookClient.UpdateBookStatus(ctx, updateStatusReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update book status: %v", err)
	}

	return &pb.ReturnBookResponse{
		Message: "Book returned successfully",
	}, nil
}
