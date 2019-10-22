package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("db_host")

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=postgres sslmode=disable password=%s", dbHost, username, password)

	db, err := gorm.Open("postgres", dbUri)
	createDbQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	db = db.Exec(createDbQuery)
	if db.Error != nil {
		fmt.Println("Unable to create DB test_db, attempting to connect assuming it exists...")
	}
	db.Close()

	dbUri = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	db, err = gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db.Debug().AutoMigrate(&Account{}, &Message{})
}

func GetDB() *gorm.DB {
	return db
}