package api

import (
	"log"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/mapper/dbmapper"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"
)

type Adapter struct {
	db   ports.PostgresSQLPort
	core ports.IDCoreFunctionsPorts
}

func NewAdapter(db ports.PostgresSQLPort, core ports.IDCoreFunctionsPorts) *Adapter {
	return &Adapter{db: db, core: core}
}

func (adapter Adapter) AddNewPersonInfo(personInfo *pb.PersonInfoRequest) error {
	personInfoModel := dbmapper.PersonInfoRequestToPersonInfoModel(personInfo)

	nationalId, _ := adapter.core.GenerateNationalID()
	personInfoModel.NationalID = db.NationalIDNumberModel{NationalID: *nationalId}

	error := adapter.db.AddNewPersonInfo(personInfoModel)
	if error != nil {
		log.Fatal("Error adding a new person")
	}
	return error
}

func (adapter Adapter) UpdatePersonInfo(personNewInfo *pb.EditPersonInfoParameters) error {
	// make a call to the DB
	personNewInfoRequest := personNewInfo.EditedPersonInfo
	personInfoModel := dbmapper.PersonInfoRequestToPersonInfoModel(personNewInfoRequest)
	personInfoModel.NationalID = *dbmapper.ProtoNationalIDNumberToNationalIDNumberModel(personNewInfo.PersonId)
	error := adapter.db.UpdatePersonInfo(personInfoModel)

	if error != nil {
		log.Fatal("Error updating person info")
	}
	return error
}

func (adapter Adapter) FindPersonInfo(nationalID *pb.NationalIDNumber) (*pb.PersonInfoResponse, error) {

	nationalIDModel := dbmapper.ProtoNationalIDNumberToNationalIDNumberModel(nationalID)
	personInfo, err := adapter.db.FindPersonInfo(nationalIDModel)

	if err != nil {
		log.Fatal("Error finding a personInfo")
		return &pb.PersonInfoResponse{}, err
	}
	result := dbmapper.PersonInfoModelToPersonInfoResponse(personInfo)
	return result, nil
}

func (adapter Adapter) GenerateNationalID() (*string, error) {
	return adapter.core.GenerateNationalID()
}
