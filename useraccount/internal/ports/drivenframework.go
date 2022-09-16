package ports

type PostgresSQLPort interface {
	OpenDBConnection()
	CloseDBConnection()
	AddNewPersonInfo()
	UpdatePersonInfo()
	FindPersonInfo()
}
