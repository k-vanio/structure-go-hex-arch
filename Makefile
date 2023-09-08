.PHONY: all
all:
	@echo "**********************************************************"
	@echo "**                    hex arch build tool                    **"
	@echo "**********************************************************"

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

protoc:
	protoc --go_out=internal/adapters/framework/left/grpc internal/adapters/framework/left/grpc/proto/number_msg.proto
	protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/left/grpc internal/adapters/framework/left/grpc/proto/arithmetic_service.proto

tidy:
	go mod tidy

.PHONY: test
test:
	go test -race -cover -v ./internal/...