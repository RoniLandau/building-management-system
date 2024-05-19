package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// DB holds the database connection
var DB *sql.DB

// Executor is a global variable that can be used as a boil.ContextExecutor
var Executor boil.ContextExecutor

// InitDB initializes the database connection
func InitDB() {
	var err error
	//TODO - fetch from .env file
	connStr := "postgres://postgres:111112@localhost:5432/building_management_system?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	// Additional initialization can go here
	Executor = boil.ContextExecutor(DB)
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}
}
