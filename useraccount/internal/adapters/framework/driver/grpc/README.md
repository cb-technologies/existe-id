How to set up protoc and the different protoc go plugins

First download the protoc compiled binary from https://github.com/protocolbuffers/protobuf/releases
Update your PATH environment variable to include the protoc executable


Then install both plugins for protoc messages and grpc services using:

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2


To generate golanguage files from the protoc compiler:

protoc --proto_path=... --go_out=  path_to_the_proto_fine

-- proto_path: should be the folder in which the protoc compiler will look into to result any import when creating different messages
--go_out: should be the file in which the go files will be created

