// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: book.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BookService_ListBooks_FullMethodName         = "/book.BookService/ListBooks"
	BookService_AddBook_FullMethodName           = "/book.BookService/AddBook"
	BookService_UpdateBook_FullMethodName        = "/book.BookService/UpdateBook"
	BookService_DeleteBook_FullMethodName        = "/book.BookService/DeleteBook"
	BookService_SearchBooks_FullMethodName       = "/book.BookService/SearchBooks"
	BookService_CheckAvailability_FullMethodName = "/book.BookService/CheckAvailability"
	BookService_UpdateBookStatus_FullMethodName  = "/book.BookService/UpdateBookStatus"
)

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookResponse, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error)
	SearchBooks(ctx context.Context, in *SearchBooksRequest, opts ...grpc.CallOption) (*SearchBooksResponse, error)
	CheckAvailability(ctx context.Context, in *CheckAvailabilityRequest, opts ...grpc.CallOption) (*CheckAvailabilityResponse, error)
	UpdateBookStatus(ctx context.Context, in *UpdateBookStatusRequest, opts ...grpc.CallOption) (*UpdateBookStatusResponse, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, BookService_ListBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddBookResponse)
	err := c.cc.Invoke(ctx, BookService_AddBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBookResponse)
	err := c.cc.Invoke(ctx, BookService_UpdateBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBookResponse)
	err := c.cc.Invoke(ctx, BookService_DeleteBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) SearchBooks(ctx context.Context, in *SearchBooksRequest, opts ...grpc.CallOption) (*SearchBooksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchBooksResponse)
	err := c.cc.Invoke(ctx, BookService_SearchBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) CheckAvailability(ctx context.Context, in *CheckAvailabilityRequest, opts ...grpc.CallOption) (*CheckAvailabilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAvailabilityResponse)
	err := c.cc.Invoke(ctx, BookService_CheckAvailability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBookStatus(ctx context.Context, in *UpdateBookStatusRequest, opts ...grpc.CallOption) (*UpdateBookStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBookStatusResponse)
	err := c.cc.Invoke(ctx, BookService_UpdateBookStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility.
type BookServiceServer interface {
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	AddBook(context.Context, *AddBookRequest) (*AddBookResponse, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error)
	SearchBooks(context.Context, *SearchBooksRequest) (*SearchBooksResponse, error)
	CheckAvailability(context.Context, *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error)
	UpdateBookStatus(context.Context, *UpdateBookStatusRequest) (*UpdateBookStatusResponse, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookServiceServer struct{}

func (UnimplementedBookServiceServer) ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBookServiceServer) AddBook(context.Context, *AddBookRequest) (*AddBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedBookServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookServiceServer) SearchBooks(context.Context, *SearchBooksRequest) (*SearchBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBooks not implemented")
}
func (UnimplementedBookServiceServer) CheckAvailability(context.Context, *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAvailability not implemented")
}
func (UnimplementedBookServiceServer) UpdateBookStatus(context.Context, *UpdateBookStatusRequest) (*UpdateBookStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookStatus not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}
func (UnimplementedBookServiceServer) testEmbeddedByValue()                     {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	// If the following call pancis, it indicates UnimplementedBookServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_ListBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_AddBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).AddBook(ctx, req.(*AddBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_UpdateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_DeleteBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_SearchBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).SearchBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_SearchBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).SearchBooks(ctx, req.(*SearchBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_CheckAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CheckAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_CheckAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CheckAvailability(ctx, req.(*CheckAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBookStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBookStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_UpdateBookStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBookStatus(ctx, req.(*UpdateBookStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _BookService_ListBooks_Handler,
		},
		{
			MethodName: "AddBook",
			Handler:    _BookService_AddBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
		{
			MethodName: "SearchBooks",
			Handler:    _BookService_SearchBooks_Handler,
		},
		{
			MethodName: "CheckAvailability",
			Handler:    _BookService_CheckAvailability_Handler,
		},
		{
			MethodName: "UpdateBookStatus",
			Handler:    _BookService_UpdateBookStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}
