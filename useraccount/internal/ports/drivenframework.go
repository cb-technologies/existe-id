package ports

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
)

type PostgresSQLPort interface {
	CloseDBConnection()
	AddNewPersonInfo(personInfo *db.PersonInfoModel) error
	UpdatePersonInfo(nationalId *db.NationalIDNumberModel)
	FindPersonInfo(nationalId *db.NationalIDNumberModel) (*db.PersonInfoModel, error)
}
