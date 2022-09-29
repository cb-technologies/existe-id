package rpc

import "github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"

type Adapter struct {
	api ports.APIPorts
}

func NewAdapter(api ports.APIPorts) *Adapter {
	return &Adapter{api: api}
}

// In the server file, we only implement the Run() methods from the GRPC port
// This will start our grpc server
func (a Adapter) Run() {

	// TODO: Implement

}
