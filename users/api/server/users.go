package server

import (
	"context"
	"go-grpc-demo/users/db"
	"go-grpc-demo/users/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	InternalApplicationError = "internal application error"
)

func (s Server) CreateUser(ctx context.Context, in *proto.CreateUser_Request) (*proto.CreateUser_Response, error) {
	user, err := db.CreateUser(s.DB, db.NewUserFormProto(in.GetUser()))
	if err != nil {
		return nil, status.Error(codes.Internal, InternalApplicationError)
	}
	resp := &proto.CreateUser_Response{User: user.ToProto()}
	return resp, nil
}

func (s Server) GetUser(ctx context.Context, in *proto.GetUser_Request) (*proto.GetUser_Response, error) {
	user, err := db.GetUser(s.DB, in.GetID())
	if err != nil {
		return nil, status.Error(codes.Internal, InternalApplicationError)
	}
	resp := &proto.GetUser_Response{User: user.ToProto()}
	return resp, nil
}
