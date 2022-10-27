package serverGrpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/constants/serviceresponse"
	"google.golang.org/grpc"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"
)

type Adapter struct {
	api ports.APIPorts
	pb.UnimplementedExistCRUDServiceServer
}

func NewAdapter(api ports.APIPorts) *Adapter {
	return &Adapter{api: api}
}

func (adapter Adapter) AddNewPerson(ctx context.Context, req *pb.PersonInfoRequest) (*pb.Response, error) {
	error := adapter.api.AddNewPerson(req)
	if error != nil {
		log.Fatal("Could not create a personInfo")
		return &pb.Response{Status: serviceresponse.FAILURE}, error
	}
	return &pb.Response{Status: serviceresponse.SUCCESS}, error
}

func (adapter Adapter) UpdatePersonInfo(ctx context.Context, req *pb.EditPersonInfoParameters) (*pb.Response, error) {

	error := adapter.api.UpdatePersonInfo(req)

	if error != nil {
		log.Fatal("Could not update person info")
		return &pb.Response{Status: serviceresponse.FAILURE}, error
	}
	return &pb.Response{Status: serviceresponse.SUCCESS}, error
}

func (adapter Adapter) FindPersonInfo(ctx context.Context, req *pb.NationalIDNumber) (*pb.PersonInfoResponse, error) {

	personInfo, err := adapter.api.FindPersonInfo(req)

	if err != nil {
		log.Fatal("Could not Find personInfo")
		return &pb.PersonInfoResponse{}, err
	}
	return personInfo, nil
}

// In the server file, we only implement the Run() methods from the GRPC port
// This will start our grpc server
var (
	port int = 8080
)

func (adapter Adapter) Run() {

	lis, err := net.Listen("tcp", "0.0.0.0:9090")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterExistCRUDServiceServer(grpcServer, adapter)

	fmt.Println("Nicolas debugging 2")

	err = grpcServer.Serve(lis)
	fmt.Println("Checking")
	if err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
