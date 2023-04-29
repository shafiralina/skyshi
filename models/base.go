package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func init() {
	//os.Setenv("db_user", "root")
	//os.Setenv("db_name", "skyshi")
	//os.Setenv("db_host", "localhost")
	//os.Setenv("db_pass", "shafira")

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost, dbName)
	conn, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
