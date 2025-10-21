package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// First try to connect to postgres to create database
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		host, user, password, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}

	// Create database if it doesn't exist
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get underlying *sql.DB: %v", err)
	}

	row := sqlDB.QueryRow(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbname))
	var exists int
	if err := row.Scan(&exists); err != nil {
		// Database doesn't exist, create it
		_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
		if err != nil {
			log.Fatalf("❌ Failed to create database: %v", err)
		}
		log.Printf("✅ Created database %s", dbname)
	}

	// Close the default postgres connection
	sqlDB.Close()

	// Connect to the actual database
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL!")
}
