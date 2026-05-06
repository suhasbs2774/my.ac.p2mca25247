package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite", "./notifications.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	// Notifications table
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS notifications (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT,
		title TEXT,
		message TEXT,
		type TEXT,
		is_read BOOLEAN DEFAULT FALSE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Users table (for email)
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
