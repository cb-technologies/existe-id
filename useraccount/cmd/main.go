package main

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/application/api"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/core/national_id_generator"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/serverGrpc"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

type Adapter2 struct {
	pb.UnimplementedExistCRUDServer
}

func main() {

	var postgresAdapter ports.PostgresSQLPort
	var appAdapter ports.APIPorts
	var coreAdapter ports.IDCoreFunctionsPorts

	// In the future we will probably pass the different variable in order to build a dsn here
	// We can get them from the environment variables
	// for now we are creating a dsn within the postgresSQL new adapter function
	postgresAdapter, _ = postgresSQL.NewAdapter()
	coreAdapter = national_id_generator.NewAdapter()
	appAdapter = api.NewAdapter(postgresAdapter, coreAdapter)
	gRPCAdapter := serverGrpc.NewAdapter(appAdapter)
	//fmt.Println("Nicolas debugging")

	grpcServer := grpc.NewServer()

	pb.RegisterExistCRUDServer(grpcServer, gRPCAdapter)
	reflection.Register(grpcServer)
	// Opening another thread to start the real exist id server

	lis, err := net.Listen("tcp", "localhost:4550")
	go func() {
		log.Fatalf("failed to serve : %v", grpcServer.Serve(lis))
	}()

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		// Enabling CORS
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	httpServerExist := &http.Server{
		Handler: grpcWebServer,
		Addr:    "localhost:4551",
	}

	log.Printf("Http server listening at %v", httpServerExist.Addr)

	err = httpServerExist.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
