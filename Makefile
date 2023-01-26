message_proto:
	protoc --proto_path=useraccount/internal/adapters/framework/driver/grpc/proto --go_out=useraccount/internal/adapters/framework/driver/grpc useraccount/internal/adapters/framework/driver/grpc/proto/message_and_service.proto

service_proto:
	protoc --go-grpc_out=useraccount/internal/adapters/framework/driver/grpc --proto_path=useraccount/internal/adapters/framework/driver/grpc/proto useraccount/internal/adapters/framework/driver/grpc/proto/message_and_service.proto

.PHONY: message_proto service_proto