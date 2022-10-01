package postgresSQL

import (
	"database/sql"
	"fmt"
	"log"

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
	dbPassword = "exist2022"
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
		log.Fatalf("dbmapper close failure: %v", err)
	}
}

func (adapter Adapter) AddNewPersonInfo(personInfo *db.PersonInfoModel) error {
	if !adapter.isDatabaseTableCreated() {
		err := adapter.createDatabaseTable()
		if err != nil {
			log.Fatalf("Could not create the database %v", err)
		}
	}
	err := adapter.db.AutoMigrate(&db.PersonInfoModel{})

	if err != nil {
		log.Fatalf("Could not migrate the database")
	}

	result := adapter.db.Create(personInfo)

	if result.Error != nil {
		log.Fatalf("Failed to create the user. error %v", result.Error)
	}

	return result.Error
}

func (adapter Adapter) UpdatePersonInfo(nationalID string) {
	// TODO: implement
	person := db.PersonInfoModel{}
	result := adapter.db.First(&person, "national_id = ?", nationalID)
	if result.Error != nil {
		//fmt.Print(err)
		log.Fatalf("Could not find the person")
	} else {
		fmt.Println(result.RowsAffected)
	}
}

func (adapter Adapter) FindPersonInfo(nationalId *db.NationalIDNumberModel) (*db.PersonInfoModel, error) {

	var result db.PersonInfoModel

	personInfo := &db.PersonInfoModel{NationalID: *nationalId}
	err := adapter.db.Where(personInfo).First(&result).Error

	if err != nil {
		log.Fatalf("Person with id %v does not exist", nationalId.NationalID)
		return &db.PersonInfoModel{}, err
	}

	return &result, nil

}

func (adapter Adapter) createDatabaseTable() error {

	err := adapter.db.Migrator().CreateTable(&db.PersonInfoModel{})
	if err != nil {
		log.Fatalf("dbmapper creation failure: %v", err)
	}
	err = adapter.db.AutoMigrate(&db.PersonInfoModel{})
	if err != nil {
		log.Fatalf("Failed to Migrate table: %v", err)
	}
	return err
}

func (adapter Adapter) isDatabaseTableCreated() bool {
	return adapter.db.Migrator().HasTable(&db.PersonInfoModel{})
}
