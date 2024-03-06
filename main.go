package main

import (
	"fmt"

	"github.com/Luis-Honorato/todo_go_api/config"
	"github.com/Luis-Honorato/todo_go_api/router"
)

func main() {
	// Initialize App require configs (db)
	err := config.InitializeConfigs()

	// On initialize error, kill application
	if err != nil {
		fmt.Printf("config initialization error %v", err)
		return
	}

	// Initialize app router
	router.InitializeRouter()
}
