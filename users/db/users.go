package db

import (
	"fmt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	err := db.Create(user).Error
	fmt.Println("User Created", user)

	return user, err
}

func GetUser(db *gorm.DB, ID uint32) (*User, error) {
	// Read
	user := &User{}
	db.First(user, ID)
	return user,nil
}

