package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./db/api.db")
	if err != nil {
		errorString := "Error initializing the database: " + err.Error()
		panic(errors.New(errorString))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	var err error
	// Create the users table
	createUsersTableStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = DB.Exec(createUsersTableStmt)
	if err != nil {
		errorString := "Error creating the users table: " + err.Error()
		panic(errors.New(errorString))
	}

	createEventsTableStmt := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(userId) REFERENCES users(id) ON DELETE CASCADE
	);
	`
	_, err = DB.Exec(createEventsTableStmt)
	if err != nil {
		errorString := "Error creating the events table: " + err.Error()
		panic(errors.New(errorString))
	}
}

func CloseDB() {
	DB.Close()
}
