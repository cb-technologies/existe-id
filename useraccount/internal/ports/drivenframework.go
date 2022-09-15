package ports

type PostgresSQLPort interface {
	OpenDBConnection()
	CloseDBConnection()
	AddNewPersonInfo()
	DeletePersonInfo()
	UpdatePersonInfo()
	FindPersonInfo()
}
