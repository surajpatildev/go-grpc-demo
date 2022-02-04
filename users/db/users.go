package db

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func FindUser(db *gorm.DB, email string) (*User, error) {
	var dbUser User
	err := db.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}
	return &dbUser, nil
}

func CheckIfEmailExists(db *gorm.DB, email string) error {
	_, err := FindUser(db, email)
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != gorm.ErrRecordNotFound {
		return status.Error(codes.AlreadyExists, "Email Already Exists.")
	}
	if err != nil {
		return status.Error(codes.Internal, "Internal Server Error.")
	}
	return nil
}

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	err := CheckIfEmailExists(db, user.Email)
	if err != nil {
		return nil, err
	}
	err = db.Create(user).Error
	fmt.Println("User Created", user)

	return user, err
}

func GetUser(db *gorm.DB, ID uint32) (*User, error) {
	// Read
	user := &User{}
	err := db.First(user, ID).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUsers(db *gorm.DB) ([]*User, error) {
	var dbUsers []*User
	err := db.Find(&dbUsers).Error
	if err != nil {
		return nil, err
	}
	if dbUsers == nil {
		dbUsers = make([]*User, 0)
	}
	return dbUsers, nil
}

func UpdateUser(db *gorm.DB, user *User) error {
	return db.Model(user).Updates(*user).Error
}
