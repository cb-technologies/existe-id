package ports

import "github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
import "github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"

type IDCoreFunctionsPorts interface {
	GenerateNationalID(personInfo *pb.PersonInfoRequest) (*string)
	VerifyPassword(hashedPassword string, passwordGiven string) error
	GetCardValidity() *db.CardValidityModel
}
