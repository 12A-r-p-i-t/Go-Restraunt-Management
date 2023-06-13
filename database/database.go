package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in database.go")
	}
	db_URL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("mysql", db_URL)
	if err != nil {
		log.Fatal("Error in establising connection to database", err)
	}
	fmt.Println("Successfully connected to DB")
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
