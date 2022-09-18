package ports

type PostgresSQLPort interface {
	CloseDBConnection()
	AddNewPersonInfo()
	UpdatePersonInfo()
	FindPersonInfo()
}
