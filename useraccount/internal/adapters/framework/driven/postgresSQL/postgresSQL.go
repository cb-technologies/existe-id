package postgresSQL

import (
	"database/sql"
	"fmt"
	"log"
)

//  Johan will probably start working here

// the adapter that will implement the interface defined in ports
type Adapter struct {
	db *sql.DB
}

const (
	host     = " localhost"
	port     = 5432
	user     = "postgres"
	password = "xx ="
	dbname   = "exist-id"
)

func NewAdapter() (*Adapter, error) {
	postgresDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", postgresDbInfo)

	if err != nil {
		panic(err) // have to know what to do in this case
	}

	defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return &Adapter{db: db}, nil
}

func (adapter Adapter) CloseDBConnection() {
	err := adapter.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}
func (adapter Adapter) AddNewPersonInfo() {
	// TODO: implement
}
func (adapter Adapter) UpdatePersonInfo() {
	// TODO: implement
}
func (adapter Adapter) FindPersonInfo() {
	// TODO: implement
}
