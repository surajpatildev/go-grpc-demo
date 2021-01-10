package db

import (
	"go-grpc-demo/users/proto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (user *User) ToProto() *proto.User {
	return &proto.User{
		ID:    uint32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
}

func NewUserFormProto(user *proto.User) *User {
	return &User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
