package postgresSQL

//  Johan will probably start working here

// the adapter that will implement the interface defined in ports
type Adapter struct {
}

func NewAdapter() *Adapter {
	// TODO: modify 
	return &Adapter{}

}

func (a Adapter) OpenDBConnection() {
	// TODO: implement
}
func (a Adapter) CloseDBConnection() {
	// TODO: implement
}
func (a Adapter) AddNewPersonInfo() {
	// TODO: implement
}
func (a Adapter) UpdatePersonInfo() {
	// TODO: implement
}
func (a Adapter) FindPersonInfo() {
	// TODO: implement
}
