package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=postgres sslmode=disable password=%s", dbHost, username, password)
	fmt.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	createDbQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	db = db.Exec(createDbQuery)
	if db.Error != nil {
		fmt.Println("Unable to create DB " + dbName + " , attempting to connect assuming it exists...")
	}
	db.Close()

	dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	db, err = gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db.Debug().AutoMigrate(&Account{}, &Message{})
}

// GetDB returns db object
func GetDB() *gorm.DB {
	return db
}
