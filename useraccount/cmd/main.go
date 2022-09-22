package main

import "fmt"
import "os"
import "github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"

func main() {
	fmt.Sprintf("%s",os.getenv())

	fmt.Println("Existe ID")
	
	postgres, err := postgresSQL.NewAdapter()

	fmt.Sprintf("The error is %s",err)
	postgres.CloseDBConnection()



}
