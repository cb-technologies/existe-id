package serverGrpc

import (
	"context"
	"log"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/constants/serviceresponse"
)

func (adapter Adapter) AddNewPersonInfo(ctx context.Context, req *pb.PersonInfoRequest) (*pb.Response, error) {
	error := adapter.api.AddNewPerson(req)
	if error != nil {
		log.Print("Could not create a personInfo")
		return &pb.Response{Status: serviceresponse.FAILURE}, error
	}
	return &pb.Response{Status: serviceresponse.SUCCESS}, error
}

func (adapter Adapter) UpdatePersonInfo(ctx context.Context, req *pb.EditPersonInfoParameters) (*pb.Response, error) {

	error := adapter.api.UpdatePersonInfo(req)

	if error != nil {
		log.Print("Could not update person info")
		return &pb.Response{Status: serviceresponse.FAILURE}, error
	}
	return &pb.Response{Status: serviceresponse.SUCCESS}, error
}

func (adapter Adapter) FindPersonInfo(ctx context.Context, req *pb.NationalIDNumber) (*pb.PersonInfoResponse, error) {

	personInfo, err := adapter.api.FindPersonInfo(req)

	if err != nil {
		log.Print("Could not Find personInfo")
		return &pb.PersonInfoResponse{}, err
	}
	return personInfo, nil
}

func (adapter Adapter) RetreiveUserBasedOnField(ctx context.Context, req *pb.RetreivePersonInfoParameters) (*pb.PersonInfoResponse, error) {

	personInfo, err := adapter.api.RetreiveUserBasedOnField(req.Names, req.DateOfBirth)

	if err != nil {
		log.Printf("[ServerGRPC] Person with name %v and date of birth %v does not exist", req.Names.Nom+" "+req.Names.Prenom, req.DateOfBirth.Month+"/"+req.DateOfBirth.Day+"/"+req.DateOfBirth.Year)
		return &pb.PersonInfoResponse{}, err
	}
	return personInfo, nil
}
