package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func init() {
	//os.Setenv("MYSQL_USER", "root")
	//os.Setenv("MYSQL_DBNAME", "skyshi")
	//os.Setenv("MYSQL_HOST", "localhost")
	//os.Setenv("MYSQL_PASSWORD", "shafira")
	//os.Setenv("MYSQL_PORT", "3306")

	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DBNAME")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	conn, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
	}
	db.AutoMigrate(&Activities{}, &Todos{})

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
