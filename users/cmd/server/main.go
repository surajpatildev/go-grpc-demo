package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"go-grpc-demo/users/api/server"
	"go-grpc-demo/users/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitialMigration for project with db.AutoMigrate
func InitialMigration() *gorm.DB {
	dsn := "host=localhost user=vs password=111 dbname=vs port=5432 sslmode=disable"
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	// Migrate the schema
	_ = con.AutoMigrate(&db.User{})
	return con
}
func main() {
	DB := InitialMigration()
	srv, err := server.NewServer(DB)
	if err != nil {
		fmt.Println("Could not create server", err)
	}


	flag.Parse()
	defer glog.Flush()

	if err := srv.Serve(); err != nil {
		glog.Fatal(err)
	}

}
