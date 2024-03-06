package handler

import (
	"github.com/Luis-Honorato/todo_go_api/config"
	"gorm.io/gorm"
)

var (
	// Declaires db as a global package variable,
	// To pass it in handlers functions
	db *gorm.DB
)

// Get db from config, to use in handlers
func InitializeHandler() {
	db = config.GetDB()

}
