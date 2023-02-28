package api

import (
	"log"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/mapper/dbmapper"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"
)

type Adapter struct {
	dbPersonInfo ports.PostgresSQLPersonInfoPort
	dbAgentInfo  ports.PostgresSQLAgentInfoPort
	core         ports.IDCoreFunctionsPorts
}

func NewAdapter(dbPersonInfo ports.PostgresSQLPersonInfoPort, dbAgentInfo ports.PostgresSQLAgentInfoPort, core ports.IDCoreFunctionsPorts) *Adapter {
	return &Adapter{dbPersonInfo: dbPersonInfo, dbAgentInfo: dbAgentInfo, core: core}
}

func (adapter Adapter) AddNewPerson(personInfo *pb.PersonInfoRequest) error {
	personInfoModel := dbmapper.PersonInfoRequestToPersonInfoModel(personInfo)
	personInfoModel.CardValidity = *adapter.core.GetCardValidity()

	nationalId := adapter.core.GenerateNationalID(personInfo)
	personInfoModel.NationalID = db.NationalIDNumberModel{NationalID: *nationalId}

	error := adapter.dbPersonInfo.AddNewPersonInfo(personInfoModel)
	if error != nil {
		log.Print("Error adding a new person")
	}
	return error
}

func (adapter Adapter) UpdatePersonInfo(personNewInfo *pb.EditPersonInfoParameters) error {
	// make a call to the DB
	personNewInfoRequest := personNewInfo.EditedPersonInfo
	personInfoModel := dbmapper.PersonInfoRequestToPersonInfoModel(personNewInfoRequest)
	personInfoModel.NationalID = *dbmapper.ProtoNationalIDNumberToNationalIDNumberModel(personNewInfo.PersonId)
	error := adapter.dbPersonInfo.UpdatePersonInfo(personInfoModel)

	if error != nil {
		log.Print("Error updating person info")
	}
	return error
}

func (adapter Adapter) FindPersonInfo(nationalID *pb.NationalIDNumber) (*pb.PersonInfoResponse, error) {

	nationalIDModel := dbmapper.ProtoNationalIDNumberToNationalIDNumberModel(nationalID)
	personInfo, err := adapter.dbPersonInfo.FindPersonInfo(nationalIDModel)

	if err != nil {
		log.Print("Error finding a personInfo")
		return &pb.PersonInfoResponse{}, err
	}
	result := dbmapper.PersonInfoModelToPersonInfoResponse(personInfo)
	return result, nil
}

func (adapter Adapter) RetreiveUserBasedOnField(names *pb.Names, dateOfBirth *pb.DateOfBirth, sex *pb.Sex) (*pb.PersonInfoResponse, error) {

	namesModel := dbmapper.ProtoNamesToDBNames(names)
	dateOfBirthModel := dbmapper.ProtoDateOfBirthToDBDateOfBirth(dateOfBirth)
	sexModel := dbmapper.protoSexToDBSex(sex)
	personInfo, err := adapter.dbPersonInfo.RetreiveUserBasedOnField(namesModel, dateOfBirthModel, sexModel)

	if err != nil {
		log.Print("Error finding a personInfo")
		return &pb.PersonInfoResponse{}, err
	}
	result := dbmapper.PersonInfoModelToPersonInfoResponse(personInfo)
	return result, nil
}

func (adapter Adapter) GenerateNationalID(personInfo *pb.PersonInfoRequest) (*string) {
	return adapter.core.GenerateNationalID(personInfo)
}

func (adapter Adapter) SignInAgent(agentSignInInfo *pb.AgentSignInInfo) error {
	agentSignInInfoModel := dbmapper.ProtoAgentSignInInfoToAgentSignInInfoModel(agentSignInInfo)

	agentSignInInfoModelResult, err := adapter.dbAgentInfo.SignInAgent(agentSignInInfoModel)
	if err != nil {
		return err
	}

	err = adapter.core.VerifyPassword(agentSignInInfoModelResult.Password, agentSignInInfoModel.Password)

	return err
}

func (adapter Adapter) SignUpAgent(agentSignUpInfo *pb.AgentInfo) error {
	agentSignUpInfoModel := dbmapper.ProtoAgentInfoToAgentInfoModel(agentSignUpInfo)

	return adapter.dbAgentInfo.SignUpAgent(agentSignUpInfoModel)
}
