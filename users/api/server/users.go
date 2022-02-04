package server

import (
	"context"
	"fmt"
	"go-grpc-demo/users/db"
	"go-grpc-demo/users/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	InternalApplicationError = "internal application error"
	NotFound                 = "Not Found"
)

func (s Server) CreateUser(ctx context.Context, in *proto.CreateUser_Request) (*proto.CreateUser_Response, error) {
	if in.User == nil {
		return nil, status.Error(codes.InvalidArgument, "User can't be null")
	}

	if in.User.Email == "" || in.User.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "Email or Password can't be null")
	}
	user, err := db.CreateUser(s.DB, db.NewUserFormProto(in.GetUser()))
	if err != nil {
		return nil, err
	}
	resp := &proto.CreateUser_Response{User: user.ToProto()}
	return resp, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *proto.UpdateUser_Request) (*proto.UpdateUser_Response, error) {

	if in.User == nil {
		return nil, status.Error(codes.InvalidArgument, "User can't be null")
	}

	if in.User.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "User ID can't be null")
	}
	if in.User.Email != "" {
		user, err := db.FindUser(s.DB, in.User.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		if user != nil && user.ID != uint(in.User.ID) {
			return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("User with email %s already exists", in.User.Email))
		}
	}
	dbUser := db.NewUserFormProto(in.User)
	dbUser.ID = uint(in.User.ID)
	err := db.UpdateUser(s.DB, dbUser)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateUser_Response{User: dbUser.ToProto()}, nil
}

func (s Server) GetUser(ctx context.Context, in *proto.GetUser_Request) (*proto.GetUser_Response, error) {
	user, err := db.GetUser(s.DB, in.GetID())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, NotFound)
		}
		return nil, status.Error(codes.Internal, InternalApplicationError)
	}
	resp := &proto.GetUser_Response{User: user.ToProto()}
	return resp, nil
}

func (s *Server) GetAllUsers(ctx context.Context, in *proto.GetAllUsers_Request) (*proto.GetAllUsers_Response, error) {
	dbUsers, err := db.GetAllUsers(s.DB)

	if err != nil {
		return nil, err
	}
	var protoUsers []*proto.User

	for _, dbUser := range dbUsers {
		protoUsers = append(protoUsers, dbUser.ToProto())
	}

	return &proto.GetAllUsers_Response{Users: protoUsers}, nil
}
