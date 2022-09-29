package rpc

import (
	"context"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/constants/serviceresponse"
	"log"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
)

func (adapter Adapter) AddNewPersonInfo(ctx context.Context, req *pb.PersonInfoRequest) (*pb.Response, error) {
	error := adapter.api.AddNewPersonInfo(req)
	if error != nil {
		log.Fatal("Could not create a personInfo")
		return &pb.Response{Status: serviceresponse.FAILURE}, error
	}
	return &pb.Response{Status: serviceresponse.SUCCESS}, error
}

func (a Adapter) UpdatePersonInfo(ctx context.Context, req *pb.EditPersonInfoParameters) (*pb.Response, error) {
	// TODO: implement
	return &pb.Response{}, nil

}
func (a Adapter) FindPersonInfo(ctx context.Context, req *pb.UUID) (*pb.PersonInfoResponse, error) {
	// TODO: implement
	return &pb.PersonInfoResponse{}, nil
}
