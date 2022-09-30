package ports

import "github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"

type APIPorts interface {
	AddNewPersonInfo(personInfo *pb.PersonInfoRequest) error
	UpdatePersonInfo(parameters *pb.EditPersonInfoParameters) error
	FindPersonInfo(nationalID *pb.NationalIDNumber) (pb.PersonInfoResponse, error)
}
