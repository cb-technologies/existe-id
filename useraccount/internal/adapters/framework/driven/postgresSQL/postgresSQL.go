package postgresSQL

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//  Johan will probably start working here

// the adapter that will implement the interface defined in ports
type Adapter struct {
	db *sql.DB
}

const (
	dbPort     = 5432
	dbUser     = "postgres"
	dbHost     = "exist-identifier.cprbzqerfjiq.us-east-1.rds.amazonaws.com"
	dbName     = "existdb"
	region     = "us-east-1"
	dbPassword = "test1234"
)

func NewAdapter() (*Adapter, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

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
