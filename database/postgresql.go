package database

import (
	"fmt"
	"log"
	"os"

	md "shorturl/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DatabaseConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		dbPassword,
	)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Database.AutoMigrate(md.Link{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	log.Println("PostgreSQL connection successful...")
}
