package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/application/api"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/core"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/serverGrpc"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	var postgresPersonInfoAdapter ports.PostgresSQLPersonInfoPort
	var postgresAgentInfoAdapter ports.PostgresSQLAgentInfoPort
	var coreAdapter ports.IDCoreFunctionsPorts
	var appAdapter ports.APIPorts

	// In the future we will probably pass the different variable in order to build a dsn here
	// We can get them from the environment variables
	// for now we are creating a dsn within the postgresSQL new adapter function
	postgresPersonInfoAdapter, _ = postgresSQL.NewPersonInfoAdapter()
	postgresAgentInfoAdapter, _ = postgresSQL.NewAgentInfoAdapter()
	coreAdapter = core.NewAdapter()
	appAdapter = api.NewAdapter(postgresPersonInfoAdapter, postgresAgentInfoAdapter, coreAdapter)

	gRPCAdapter := serverGrpc.NewAdapter(appAdapter)
	//fmt.Println("Nicolas debugging")

	// Runn the HTTP server in another go routine
	go runGatewayServer(gRPCAdapter)
	// Run the GRPC server in the main routine
	runGrpcServer(gRPCAdapter)

}

func runGrpcServer(grpcAdapter *serverGrpc.Adapter) {
	grpcServer := grpc.NewServer()

	pb.RegisterExistCRUDServer(grpcServer, grpcAdapter)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "localhost:4550")
	if err != nil {
		log.Printf("failed to serve : %v", grpcServer.Serve(listener))
	}
	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)
	httpServerExist := &http.Server{
		Handler: grpcWebServer,
		Addr:    "localhost:4500",
	}
	log.Printf("GRPC server listening at %v", httpServerExist.Addr)

	err = httpServerExist.ListenAndServe()
	if err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func runGatewayServer(grpcAdapter *serverGrpc.Adapter) {

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterExistCRUDHandlerServer(ctx, grpcMux, grpcAdapter)
	if err != nil {
		log.Printf("cannot register http handler server: %v", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Printf("failed to serve : %v", listener)
	}

	log.Printf("Http server listening at %v", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Printf("failed to serve: %v", err)
	}

}
