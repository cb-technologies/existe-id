package postgreSQL

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/lib/pq"
)

func main(){
	var dbName string = ""
	var dbUser string = ""
	var dbHost string = ""
	var dbPort int = 5432
	var dbEndpoint string = fmt.Sprintf("%s:%d", dbHost, dbPort)
	var region string = "us-east-1"
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
    	panic("configuration error: " + err.Error())
	}
	authenticationToken, err := auth.BuildAuthToken(
    	context.TODO(), dbEndpoint, region, dbUser, cfg.Credentials)
    if err != nil {
	    panic("failed to create authentication token: " + err.Error())
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
        dbHost, dbPort, dbUser, authenticationToken, dbName,
    )

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }

}
