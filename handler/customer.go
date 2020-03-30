package handler

import (
	"context"
	"fmt"

	"github.com/MarthaSutopo13/go-grpc-protobuf/proto"
)

func CreateCustomer(d proto.CustomersClient) {
	req := &proto.CreateCustomerRequest{
		Customer: &proto.Customer{
			Firstname: "Martha",
			Lastname:  "Sutopo",
			Email:     "martha.sutopo@gmail.com",
		},
	}

	res, err := d.CreateCustomer(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
