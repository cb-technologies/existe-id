package main

import (
	"fmt"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/application/api"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/core/national_id_generator"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/serverGrpc"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"
)

func main() {

	var postgresAdapter ports.PostgresSQLPort
	var appAdapter ports.APIPorts
	var coreAdapter ports.IDCoreFunctionsPorts
	var gRPCAdapter ports.GRPCPort

	// In the future we will probably pass the different variable in order to build a dsn here
	// We can get them from the environment variables
	// for now we are creating a dsn within the postgresSQL new adapter function
	postgresAdapter, _ = postgresSQL.NewAdapter()
	coreAdapter = national_id_generator.NewAdapter()
	appAdapter = api.NewAdapter(postgresAdapter, coreAdapter)
	gRPCAdapter = serverGrpc.NewAdapter(appAdapter)
	fmt.Println("Nicolas debugging")
	gRPCAdapter.Run()

}
