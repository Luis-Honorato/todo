package handler

import (
	"github.com/Luis-Honorato/todo_go_api/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetDB()

}
