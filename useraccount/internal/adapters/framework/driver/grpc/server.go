package rpc


type Adapter struct {

}

func NewAdapter() *Adapter {
	return &Adapter{}
}

// In the server file, we only implement the Run() methods from the GRPC port
// This will start our grpc server
func (a Adapter) Run() {

	// TODO: Implement

}