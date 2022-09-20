package ports

import "github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"

type APIPorts interface {
	AddNewPersonInfo(personInfo pb.PersonInfo) error
	UpdatePersonInfo(parameters pb.EditPersonInfoParameters) error
	FindPersonInfo(uuid pb.UUID) (pb.PersonInfo, error)
}
