package postgresSQL

import (
	"database/sql"
	"log"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PersonInfoAdapter struct {
	db *gorm.DB
}

func NewPersonInfoAdapter() (*PersonInfoAdapter, error) {

	dsn := makeDSN()
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

	return &PersonInfoAdapter{db: db}, nil
}

func (adapter PersonInfoAdapter) CloseDBConnection() {
	db, _ := adapter.db.DB()

	err := db.Close()
	if err != nil {
		log.Fatalf("dbmapper close failure: %v", err)
	}
}

func (adapter PersonInfoAdapter) AddNewPersonInfo(personInfo *db.PersonInfoModel) error {
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

func (adapter PersonInfoAdapter) UpdatePersonInfo(personNewInfo *db.PersonInfoModel) error {

	nationalId := personNewInfo.NationalID

	person, findErr := adapter.FindPersonInfo(&nationalId)

	if findErr != nil {
		log.Fatalf("Person with id %v does not exist", nationalId.NationalID)
		return findErr
	}

	err := adapter.db.Model(db.PersonInfoModel{}).Where(&person).Updates(personNewInfo).Error

	if err != nil {
		log.Fatalf("Failed to update the user. error %v", err)
	}

	return nil
}

func (adapter PersonInfoAdapter) FindPersonInfo(nationalId *db.NationalIDNumberModel) (*db.PersonInfoModel, error) {

	var result db.PersonInfoModel

	personInfo := &db.PersonInfoModel{NationalID: *nationalId}
	err := adapter.db.Where(personInfo).First(&result).Error

	if err != nil {
		log.Fatalf("Person with id %v does not exist", nationalId.NationalID)
		return &db.PersonInfoModel{}, err
	}
	return &result, nil

}

func (adapter PersonInfoAdapter) createDatabaseTable() error {

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

func (adapter PersonInfoAdapter) isDatabaseTableCreated() bool {
	return adapter.db.Migrator().HasTable(&db.PersonInfoModel{})
}
