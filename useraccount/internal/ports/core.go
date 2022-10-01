package ports

type IDCoreFunctionsPorts interface {
	GenerateNationalID() (*string, error)
}
