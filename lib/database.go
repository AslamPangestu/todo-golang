package lib

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	var USERNAME = os.Getenv("DB_USERNAME")
	var PASSWORD = os.Getenv("DB_PASSWORD")
	var HOST = os.Getenv("DB_HOST")
	var PORT = os.Getenv("DB_PORT")
	var DBNAME = os.Getenv("DB_NAME")
	var SSL_MODE = os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", HOST, USERNAME, PASSWORD, DBNAME, PORT, SSL_MODE)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB Connection Failure, %v\n", err.Error())
	}
	fmt.Println("DB Connected")
	return db
}
