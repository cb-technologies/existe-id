package api

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
)

type DatabaseAction interface {
	AddNewPersonInfo(personInfo *pb.PersonInfoRequest) error
	UpdatePersonInfo(parameters *pb.EditPersonInfoParameters) error
	FindPersonInfo(nationalID *pb.NationalIDNumber) (pb.PersonInfoResponse, error)
}

type CoreFunctions interface {
	GenerateNationalID(model *db.PersonInfoModel)
}
