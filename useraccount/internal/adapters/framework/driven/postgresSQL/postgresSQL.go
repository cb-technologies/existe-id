package postgresSQL

import (
	"database/sql"
	"fmt"
	"log"
	"context"
	"database/sql"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/lib/pq"

)

//  Johan will probably start working here

// the adapter that will implement the interface defined in ports
type Adapter struct {
	db *sql.DB
}

const (
	dbPort      = 5432
	dbUser      = "postgres"
	dbHost 		= ""
	dbName      = "existdb"
	region		= "us-east-1"
)

func NewAdapter() (*Adapter, error) {
	var dbEndpoint string = fmt.Sprintf("%s:%d", dbHost, dbPort)
	// postgresDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

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
