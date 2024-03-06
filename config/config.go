package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	// Declaires db as a global package variable,
	// To pass it in GetDB func
	db *gorm.DB
)

// Instanciate used dependecies
func InitializeConfigs() error {
	// Declaires error. (Cannot declaires while assign, because
	// db var is alredy declaired)
	var err error

	// Assing DB with initializeDB method
	db, err = InitializeDB()

	// When ocurres error initializating DB, returns error
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)

	}

	// If reach here. don't occurred any error and returns nil
	return nil
}

// Function to returns db, used in handler
func GetDB() *gorm.DB {
	return db
}
