package client

import (
	"go-grpc-demo/users/proto"
	"google.golang.org/grpc"
	"log"
)

// Client provides services to communicating with the server
type Client struct {
	conn *grpc.ClientConn

	UsersImpl proto.UserServiceClient
}

// NewClient creates a new client from configuration
func NewClient() (*Client, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Failed to create client,", err)
	}
	return &Client{
		conn: conn,
		UsersImpl: proto.NewUserServiceClient(conn),
	}, nil
}
