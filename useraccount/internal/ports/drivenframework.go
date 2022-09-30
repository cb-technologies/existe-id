package ports

import "github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"

type PostgresSQLPort interface {
	CloseDBConnection()
	AddNewPersonInfo(personInfo *pb.PersonInfoRequest) error
	UpdatePersonInfo()
	FindPersonInfo(nationalId *pb.NationalIDNumber) (*pb.PersonInfoResponse, error)
}
