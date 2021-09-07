package services

import (
	"context"
	"fmt"

	"github.com/antunesgabriel/grpc/gen"
)

type UserService struct {
	gen.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, req *gen.UserCreateDTO) (*gen.User, error) {

	fmt.Println(req.Name)

	return &gen.User{
		Id: "RANDOM ID",	
		Name: req.GetName(),
		Email: req.GetEmail(),
		Age: req.GetAge(),	
	}, nil
}

