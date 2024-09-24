PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GRPC_GO := $(shell go env GOPATH)/bin/protoc-gen-go-grpc

build-proto:
	@protoc --go_out=. --go-grpc_out=. --proto_path=./book-service/proto book-service/proto/book.proto
	@protoc --go_out=. --go-grpc_out=. --proto_path=./user-service/proto user-service/proto/user.proto
	@protoc --go_out=. --go-grpc_out=. --proto_path=./borrow-service/proto borrow-service/proto/borrow.proto
	@protoc --go_out=. --go-grpc_out=. --proto_path=./auth-service/proto auth-service/proto/auth.proto