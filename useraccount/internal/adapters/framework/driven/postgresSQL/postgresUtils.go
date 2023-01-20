package postgresSQL

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
    "os"
    "encoding/json"
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
    //sess, err := session.Must(session.NewSession(&aws.Config{
      //  Region:      aws.String("us-east-1"),
        //Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
    //}))
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String("us-east-1"),
        Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
    })


    if err != nil {
    fmt.Println("Error creating session: ", err)
} else {
    fmt.Println("Session created successfully")
}

    // Then we create a new Secrets Manager client
    client := secretsmanager.New(sess)

    // Use the client to retrieve the secret value
    result, err := client.GetSecretValue(&secretsmanager.GetSecretValueInput{
        SecretId: aws.String("arn:aws:secretsmanager:us-east-1:456807214525:secret:rds!cluster-efc93003-9c72-432e-8b3c-613bf8a1e11c-8trB0Q"),
    })
    if err != nil {
        fmt.Println("here is the error",err)
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
