package main

import (
	"context"
	"fmt"
	"net"

	"github.com/MarthaSutopo13/go-grpc-protobuf/proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {

	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	serv := grpc.NewServer()
	proto.RegisterCustomersServer(serv, &server{})

	if err := serv.Serve(l); err != nil {
		panic(err)
	}
}

func (s *server) CreateCustomer(ctx context.Context, req *proto.CreateCustomerRequest) (*proto.CreateCustomerResponse, error) {
	firstname := req.GetCustomer().Firstname
	lastname := req.GetCustomer().Lastname
	email := req.GetCustomer().Email

	fmt.Println("context: ", ctx.Done())
	res := &proto.CreateCustomerResponse{
		Customer: &proto.Customer{
			Firstname: firstname,
			Lastname:  lastname,
			Email:     email,
		},
	}

	fmt.Println("aaa", res)
	return res, nil
}
