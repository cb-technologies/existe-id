package main

import "fmt"
import "github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"

func main() {

	fmt.Println("Existe ID")
	
	postgres, err := postgresSQL.NewAdapter()

	fmt.Sprintf("The error is %s",err)
	postgres.CloseDBConnection()



}
