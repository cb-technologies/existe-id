package postgresSQL

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
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

	postgresdb, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = postgresdb.Ping()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: postgresdb,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Adapter{db: db}, nil
}

func (adapter Adapter) CloseDBConnection() {
	db, _ := adapter.db.DB()

	err := db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (adapter Adapter) AddNewPersonInfo(personInfo *pb.PersonInfoRequest) {

	// personInformationRequest := mapper.protoPersonInfoRequestToDBPersonInfoRequest()

	if adapter.isDatabaseTableCreated() {
		// add directly
	}
	// create
	err := adapter.createDatabaseTable()

	if err != nil {
		// do something
	}
	// add user
	//adapter.db.Create(personInformation)

}

func (adapter Adapter) UpdatePersonInfo() {
	// TODO: implement
}
func (adapter Adapter) FindPersonInfo() {
	// TODO: implement
}

func (adapter Adapter) createDatabaseTable() error {

	err := adapter.db.Migrator().CreateTable(&db.PersonInfoRequest{})
	if err != nil {
		log.Fatalf("db creation failure: %v", err)
	}
	err = adapter.db.AutoMigrate(&db.PersonInfoRequest{})
	if err != nil {
		log.Fatalf("Failed to Migrate table: %v", err)
	}
	return err
}

func (adapter Adapter) isDatabaseTableCreated() bool {
	return adapter.db.Migrator().HasTable(&db.PersonInfoRequest{})
}
