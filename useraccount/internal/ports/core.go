package ports

type IDCoreFunctionsPorts interface {
	GenerateNationalID() (*string, error)
	VerifyPassword(hashedPassword string, passwordGiven string) error
}
