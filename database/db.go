package database

import (
	"log"
	"os"
	"path/filepath"
	"product-api/models"

	"github.com/glebarez/sqlite" // Pure Go SQLite driver for GORM, no CGO needed
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	// Create database directory if it doesn't exist
	dbDir := "./data"
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		err = os.MkdirAll(dbDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create database directory: %v", err)
		}
	}

	// Open the SQLite database file
	dbPath := filepath.Join(dbDir, "products.db")

	// Configure GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // For more detailed logging
	}

	// Open the database connection with pure Go driver
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migrate the schema
	migrateDB()

	log.Printf("Connected to SQLite database at %s", dbPath)
}

// migrateDB automatically creates/updates database tables based on models
func migrateDB() {
	if err := DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Error getting DB instance: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
		return
	}
	log.Println("Database connection closed")
} 