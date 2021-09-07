package main

import (
	"context"
	"fmt"
	"log"

	"github.com/antunesgabriel/grpc/gen"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure)

	if err != nil {
		log.Fatalf("Could not connect client: %v", err)
	}

	defer connection.Close()

	client := gen.NewUserServiceClient(connection)

	AddUser(client)
}


func AddUser(client gen.UserServiceClient) {
	
	req := &gen.UserCreateDTO{
		Name: "Gabriel Antunes",
		Email: "gabriel@antunes.com",
		Age: 24,
	}

	user, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Fail to register user in grpc connection: %v", err)
	}

	fmt.Println(user)
}