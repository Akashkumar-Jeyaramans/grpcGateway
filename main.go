package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	gen "github.com/Akashkumar-Jeyaramans/grpcGateway/v1/commands"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type snmpgrpcServer struct {
	gen.UnimplementedGreeterServer
}

func (s *snmpgrpcServer) SayHello(ctx context.Context, req *gen.HelloRequest) (*gen.HelloReply, error) {
	return &gen.HelloReply{
		Message: fmt.Sprintf("hello %s ", req.Name),
	}, nil
}

func main() {
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux()
	// setting up a dail up for gRPC service by specifying endpoint/target url
	//err := gen.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, "localhost:50051", []grpc.DialOption{grpc.WithInsecure()})
	err := gen.RegisterGreeterHandlerServer(context.Background(), mux, &snmpgrpcServer{})
	if err != nil {
		log.Fatal(err)
	}
	// Creating a normal HTTP server
	server := http.Server{
		Handler: mux,
	}
	// creating a listener for server
	l, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}
	// start server
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
