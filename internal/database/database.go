package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s database=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		return db, err
	}
	if err := db.DB().Ping(); err != nil {
		return db, err
	}
	return db, nil

}
