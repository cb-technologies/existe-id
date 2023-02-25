proto:
	protoc --proto_path=useraccount/internal/adapters/framework/driver/grpc/proto \
			--go-grpc_out=useraccount/internal/adapters/framework/driver/grpc \
			--go_out=useraccount/internal/adapters/framework/driver/grpc \
			--grpc-gateway_out=useraccount/internal/adapters/framework/driver/grpc/pb \
            --grpc-gateway_opt=paths=source_relative \
			useraccount/internal/adapters/framework/driver/grpc/proto/message_and_service.proto
server:
	go run useraccount/cmd/main.go

.PHONY: proto server