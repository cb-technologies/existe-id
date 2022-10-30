package ports

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
)

type PostgresSQLPersonInfoPort interface {
	CloseDBConnection()
	AddNewPersonInfo(personInfo *db.PersonInfoModel) error
	UpdatePersonInfo(personNewInfo *db.PersonInfoModel) error
	FindPersonInfo(nationalId *db.NationalIDNumberModel) (*db.PersonInfoModel, error)
}

type PostgresSQLAgentInfoPort interface {
	SignInAgent(agentSignInInfo *db.AgentSignInModel) (*db.AgentSignInModel, error)
	SignUpAgent(agentSignUpInfo *db.AgentInfoModel) error
	CloseDBConnection()
}
