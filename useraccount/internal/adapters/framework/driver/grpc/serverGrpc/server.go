package serverGrpc

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"
)

type Adapter struct {
	pb.UnimplementedExistCRUDServer
	api ports.APIPorts
}

func NewAdapter(api ports.APIPorts) *Adapter {
	return &Adapter{api: api}
}
