package config

import (
	"fmt"
	"os"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create DB file and open DataBase.
//
// Returns a pointer of Gorm Data Base
func InitializeDB() (*gorm.DB, error) {
	// Set path to alocate db
	dbPath := "./db/main.db"

	// Recovery the status of db file
	_, err := os.Stat(dbPath)

	// If file is not created
	if os.IsNotExist(err) {
		fmt.Println("Database file not found, creating...")

		// Create a file named db, to storage database
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		// Create a file on given dbPath
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		// Close the file to keep logic to open db
		file.Close()

	}

	// Once file is created

	// Get db from gorm lib, using default configs
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		fmt.Print("error opening db")
		return nil, err
	}

	// Migrate tables from db
	if err := db.AutoMigrate(schemas.Todo{}); err != nil {
		fmt.Printf("sqlite automigration error: %v", err)
		return nil, err
	}

	// Returns db
	return db, nil
}
