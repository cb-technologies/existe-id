package postgresSQL

import (
	"database/sql"
	"log"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AgentInfoAdapter struct {
	db *gorm.DB
}

func NewAgentInfoAdapter() (*AgentInfoAdapter, error) {
	dsn := makeDSN()
	postgresdb, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: postgresdb,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &AgentInfoAdapter{db: db}, nil
}

func (adapter AgentInfoAdapter) createDatabaseTable() error {

	err := adapter.db.Migrator().CreateTable(&db.AgentInfoModel{})
	if err != nil {
		log.Fatalf("dbmapper createion failure: %v", err)
	}
	err = adapter.db.AutoMigrate(&db.AgentInfoModel{})
	if err != nil {
		log.Fatalf("Failed to Migrate table: %v", err)
	}

	return err
}

func (adapter AgentInfoAdapter) isDatabaseTableCreated() bool {
	return adapter.db.Migrator().HasTable(&db.AgentInfoModel{})
}

func (adapter AgentInfoAdapter) CloseDBConnection() {
	db, _ := adapter.db.DB()

	err := db.Close()
	if err != nil {
		log.Fatalf("dbmapper close failure: %v", err)
	}
}

func (adapter AgentInfoAdapter) SignUpAgent(agentSignUpInfo *db.AgentInfoModel) error {
	if !adapter.isDatabaseTableCreated() {
		err := adapter.createDatabaseTable()
		if err != nil {
			log.Fatalf("Could not create the database %v", err)
		}
	}
	err := adapter.db.AutoMigrate(&db.AgentInfoModel{})

	if err != nil {
		log.Fatalf("Could not migrate the database")
	}

	result := adapter.db.Create(agentSignUpInfo)

	if result.Error != nil {
		log.Fatalf("Failed to create the user. error %v", result.Error)
	}

	return result.Error
}

// func (adapter AgentInfoAdapter) UpdateAgentInfo() error {
// }
func (adapter AgentInfoAdapter) SignInAgent(agentSignInInfo *db.AgentSignInModel) (*db.AgentSignInModel, error) {
	var result db.AgentInfoModel

	agentInfoModel := &db.AgentInfoModel{Email: agentSignInInfo.Email}
	err := adapter.db.Where(agentInfoModel).First(&result).Error

	if err != nil {
		return &db.AgentSignInModel{}, err
	}

	return &db.AgentSignInModel{Email: result.Email, Password: result.Password}, err
}
