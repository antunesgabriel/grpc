package main

import (
	"log"
	"net"

	"github.com/antunesgabriel/grpc/gen"
	"github.com/antunesgabriel/grpc/services"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()

	gen.RegisterUserServiceServer(grpcServer, &services.UserService{})

	

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Could not connect server: %v", err)
	}
}