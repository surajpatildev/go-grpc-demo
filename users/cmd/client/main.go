package main

import (
	"context"
	"fmt"
	"go-grpc-demo/users/api/client"
	"go-grpc-demo/users/proto"
)

func main() {
	c, err := client.NewClient()
	if err != nil {
		panic("Unable to get client")
	}
	//res, err := c.UsersImpl.CreateUser(context.Background(), &proto.CreateUser_Request{User: &proto.User{Email: "test2@email.com", Name: "Test 2"}})
	//if err != nil {
	//	fmt.Println("Failed to create user,", err)
	//}
	//fmt.Println("Created User-", res.GetUser())

	res, err := c.UsersImpl.GetUser(context.Background(), &proto.GetUser_Request{ID: 2})
	if err != nil {
		fmt.Println("Failed to get user,", err)
	}
	fmt.Println("User - ", res.GetUser())

}
