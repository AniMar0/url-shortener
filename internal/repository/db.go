package repository

import (
	"database/sql"
	"fmt"
	"log"
)

// Initialize the database connection and create necessary tables
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	createTable(db)

	return db
}

// createTable creates the necessary tables for the application if they do not already exist
func createTable(db *sql.DB) {
	_, _ = db.Exec("PRAGMA foreign_keys = ON;")

	query := []string{`CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`,
		`
CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    original_url TEXT NOT NULL,
    short_code TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`,
	}

	for index, q := range query {
		_, err := db.Exec(q)
		if err != nil {
			log.Fatal("Failed to create table at index", index, ":", err)
		}
	}

	fmt.Println("Database schema is ready.")
}
