package main

import (
	"fmt"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"
)

// This file is responsible for orchestrating the startup of our application
func main() {

	fmt.Println("Existe ID")

	postgres, err := postgresSQL.NewAdapter()

	if err != nil {
		fmt.Println("The error is ", err)
	}

	fmt.Println("Connection succesful!")
	postgres.PlayField()
	postgres.CloseDBConnection()

}
