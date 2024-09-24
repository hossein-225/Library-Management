// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: borrow.proto

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
	BorrowService_BorrowBook_FullMethodName = "/borrow.BorrowService/BorrowBook"
	BorrowService_ReturnBook_FullMethodName = "/borrow.BorrowService/ReturnBook"
)

// BorrowServiceClient is the client API for BorrowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BorrowServiceClient interface {
	BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookResponse, error)
	ReturnBook(ctx context.Context, in *ReturnBookRequest, opts ...grpc.CallOption) (*ReturnBookResponse, error)
}

type borrowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBorrowServiceClient(cc grpc.ClientConnInterface) BorrowServiceClient {
	return &borrowServiceClient{cc}
}

func (c *borrowServiceClient) BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BorrowBookResponse)
	err := c.cc.Invoke(ctx, BorrowService_BorrowBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowServiceClient) ReturnBook(ctx context.Context, in *ReturnBookRequest, opts ...grpc.CallOption) (*ReturnBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReturnBookResponse)
	err := c.cc.Invoke(ctx, BorrowService_ReturnBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BorrowServiceServer is the server API for BorrowService service.
// All implementations must embed UnimplementedBorrowServiceServer
// for forward compatibility.
type BorrowServiceServer interface {
	BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookResponse, error)
	ReturnBook(context.Context, *ReturnBookRequest) (*ReturnBookResponse, error)
	mustEmbedUnimplementedBorrowServiceServer()
}

// UnimplementedBorrowServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBorrowServiceServer struct{}

func (UnimplementedBorrowServiceServer) BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BorrowBook not implemented")
}
func (UnimplementedBorrowServiceServer) ReturnBook(context.Context, *ReturnBookRequest) (*ReturnBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnBook not implemented")
}
func (UnimplementedBorrowServiceServer) mustEmbedUnimplementedBorrowServiceServer() {}
func (UnimplementedBorrowServiceServer) testEmbeddedByValue()                       {}

// UnsafeBorrowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BorrowServiceServer will
// result in compilation errors.
type UnsafeBorrowServiceServer interface {
	mustEmbedUnimplementedBorrowServiceServer()
}

func RegisterBorrowServiceServer(s grpc.ServiceRegistrar, srv BorrowServiceServer) {
	// If the following call pancis, it indicates UnimplementedBorrowServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BorrowService_ServiceDesc, srv)
}

func _BorrowService_BorrowBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowServiceServer).BorrowBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BorrowService_BorrowBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowServiceServer).BorrowBook(ctx, req.(*BorrowBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BorrowService_ReturnBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowServiceServer).ReturnBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BorrowService_ReturnBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowServiceServer).ReturnBook(ctx, req.(*ReturnBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BorrowService_ServiceDesc is the grpc.ServiceDesc for BorrowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BorrowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "borrow.BorrowService",
	HandlerType: (*BorrowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BorrowBook",
			Handler:    _BorrowService_BorrowBook_Handler,
		},
		{
			MethodName: "ReturnBook",
			Handler:    _BorrowService_ReturnBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "borrow.proto",
}
