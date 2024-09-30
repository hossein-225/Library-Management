PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GRPC_GO := $(shell go env GOPATH)/bin/protoc-gen-go-grpc

build-grpc-proto:
	@protoc --go_out=. --go-grpc_out=. --proto_path=proto proto/borrow.proto

build-mock:
	@find ./ -type f -name "*_grpc.pb.go" -exec sh -c 'echo {} && mkdir -p "$$(dirname {})/mock" && mockgen -source={} -destination="$$(dirname {})/mock/grpc.go"' \;

build-proto: build-grpc-proto build-mock

build-doc:
	@swag init -g cmd/main.go