package ports

import "github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"

type IDCoreFunctionsPorts interface {
	GenerateNationalID() (*string, error)
	VerifyPassword(hashedPassword string, passwordGiven string) error
	GetCurrentDate() *db.CardValidityModel
}
