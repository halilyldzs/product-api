package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite" // Pure Go SQLite driver, no CGO needed
)

var DB *sql.DB

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
	db, err := sql.Open("sqlite", dbPath) // Driver name is "sqlite" (not "sqlite3")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create products table if it doesn't exist
	createTable(db)

	DB = db
	log.Printf("Connected to SQLite database at %s", dbPath)
}

// createTable creates the products table if it doesn't exist
func createTable(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		quantity INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		sku TEXT
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create products table: %v", err)
	}
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
} 