package postgresSQL

import (
    "database/sql"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
    "github.com/lib/pq"
    "os"
 "fmt"
)

const (
	dbPort     = 5432
	dbUser     = "postgres"
	dbHost     = "exist-identifier.cprbzqerfjiq.us-east-1.rds.amazonaws.com"
	dbName     = "existdb"
	region     = "us-east-1"
	dbPassword = "existnaruto5"
)

func makeDSN() string {

    // Let's start by getting the access key and secret key from environment variables
    accessKey := os.Getenv("AWS_ACCESS_KEY")
    secretKey := os.Getenv("AWS_SECRET_KEY")

    // Then we create a new AWS session
    sess := session.Must(session.NewSession(&aws.Config{
        Region:      aws.String("us-west-2"),
        Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
    }))

    / Then we create a new Secrets Manager client
    client := secretsmanager.New(sess)

    // Use the client to retrieve the secret value
    result, err := client.GetSecretValue(&secretsmanager.GetSecretValueInput{
        SecretId: aws.String("my-secret"),
    })
    if err != nil {
        fmt.Println(err)
    }

    // Now we have to parse the secret value as JSON to get the database credentials
    var secret struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        User     string `json:"user"`

        Password string `json:"password"`
        Database string `json:"database"`
    }
    json.Unmarshal([]byte(*result.SecretString), &secret)

    // Use the credentials to connect to the database

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		secret.Host, secret.Port, secret.User, secret.Password, secret.Database,
	)
}
