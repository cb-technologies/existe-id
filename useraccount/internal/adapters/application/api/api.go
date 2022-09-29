package api

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"
	"log"
)

type Adapter struct {
	db ports.PostgresSQLPort
}

func NewAdapter(db ports.PostgresSQLPort) *Adapter {
	return &Adapter{db: db}
}

func (adapter Adapter) AddNewPersonInfo(personInfo *pb.PersonInfoRequest) error {
	error := adapter.db.AddNewPersonInfo(personInfo)
	if error != nil {
		log.Fatal("Error adding a new person")
	}
	return error
}

func (adapter Adapter) UpdatePersonInfo(parameters *pb.EditPersonInfoParameters) error {
	// make a call to the DB

	return nil
}

func (adapter Adapter) FindPersonInfo(uuid *pb.UUID) (pb.PersonInfoResponse, error) {
	error

	return pb.PersonInfoResponse{}, nil
}
