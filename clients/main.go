package main

import (
	"fmt"

	"github.com/MarthaSutopo13/go-grpc-protobuf/handler"
	"github.com/MarthaSutopo13/go-grpc-protobuf/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	c := proto.NewCustomersClient(conn)
	fmt.Printf("Created client customer: %v\n", c)

	handler.CreateCustomer(c)
}
