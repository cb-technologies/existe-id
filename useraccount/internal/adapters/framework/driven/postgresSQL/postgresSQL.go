package postgresSQL

import (
	"database/sql"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	_ "github.com/lib/pq"

)

//  Johan will probably start working here

// the adapter that will implement the interface defined in ports
type Adapter struct {
	db *sql.DB
}

const (
	dbPort      = 5432
	dbUser      = "postgres"
	dbHost 		= "exist-identifier.cprbzqerfjiq.us-east-1.rds.amazonaws.com"
	dbName      = "existdb"
	region		= "us-east-1"
)

func NewAdapter() (*Adapter, error) {
	var dbEndpoint string = fmt.Sprintf("%s:%d", dbHost, dbPort)
	creds := credentials.NewEnvCredentials()
    authToken, err := rdsutils.BuildAuthToken(dbEndpoint, region, dbUser, creds)
    if err != nil {
        panic(err)
    }

    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
        dbHost, dbPort, dbUser, authToken, dbName,
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
