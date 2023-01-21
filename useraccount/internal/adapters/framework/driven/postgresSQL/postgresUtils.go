package postgresSQL

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

const (
	dbPort = 5432
	dbHost = "exist-identifier.cprbzqerfjiq.us-east-1.rds.amazonaws.com"
	dbName = "existdb"
	region = "us-east-1"
)

func makeDSN() string {

	creds, err := getCredentials("rds!cluster-efc93003-9c72-432e-8b3c-613bf8a1e11c")
	if err != nil {
		log.Fatalf("Error getting Postgres credentials: %v", err)
	}

	// Connect to the Postgres server
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbHost, dbPort, creds.Username, creds.Password, dbName)

	return connStr
}

type postgresCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getCredentials(secretName string) (postgresCredentials, error) {
	var creds postgresCredentials

	// accessKey := os.Getenv("AWS_ACCESS_KEY")
	// secretKey := os.Getenv("AWS_SECRET_KEY")

	// Create a new session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials("AKIAWUW6U5W6Q6G2BUGV", "WXVwLgRGs2QAzrQch2gKjRTJ3NSPrHA7w+sVU24J", ""),
		// Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return creds, fmt.Errorf("Error creating session: %v", err)
	}

	// Create a new Secrets Manager client
	client := secretsmanager.New(sess)

	// Build the request to retrieve the secret
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	// Send the request to retrieve the secret
	result, err := client.GetSecretValue(input)
	if err != nil {
		return creds, fmt.Errorf("Error getting secret value: %v", err)
	}

	// Unmarshal the secret value
	if err := json.Unmarshal([]byte(*result.SecretString), &creds); err != nil {
		return creds, fmt.Errorf("Error unmarshalling secret value: %v", err)
	}

	return creds, nil
}
