package rpc

import (
	"context"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
)

func (a Adapter) AddNewPersonInfo(ctx context.Context, req *pb.PersonInfoRequest) (*pb.Response, error) {
	// TODO: implement
	return &pb.Response{}, nil
}
func (a Adapter) UpdatePersonInfo(ctx context.Context, req *pb.EditPersonInfoParameters) (*pb.Response, error) {
	// TODO: implement
	return &pb.Response{}, nil

}
func (a Adapter) FindPersonInfo(ctx context.Context, req *pb.UUID) (*pb.PersonInfoResponse, error) {
	// TODO: implement
	return &pb.PersonInfoResponse{}, nil
}
