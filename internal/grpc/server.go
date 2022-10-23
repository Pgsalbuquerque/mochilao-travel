package grpc

import (
	"context"
	"log"
	"net"
	"net/http"

	"mochilao-travel/internal/config"
	gen "mochilao-travel/internal/grpc/gen/go"
	grpcservice "mochilao-travel/internal/grpc/service"
	"mochilao-travel/internal/travel"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func StartGrpcServer(travels *travel.Travel) error {

	grpcServer := grpc.NewServer()

	travelServer := grpcservice.NewTravelServer(travels)

	gen.RegisterTravelServer(grpcServer, travelServer)

	port := config.Get().GRPCPort

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Grpc Server running at port: %v", port)

	grpc_Error := grpcServer.Serve(listener)
	if grpc_Error != nil {
		log.Fatal(grpc_Error)
		return grpc_Error
	}

	return nil
}

func StartGrpcGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	gen.RegisterTravelHandlerFromEndpoint(ctx, mux, config.Get().GRPCPort, []grpc.DialOption{grpc.WithInsecure()})

	port := config.Get().GRPCGatewayPort
	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("Grpc Gateway running at port: %v", port)

	// start server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
