package router

import "github.com/gin-gonic/gin"

// Gets instance of Gin and open routes ind :8080 ports
func InitializeRouter() {
	// Initialize router with default configs
	router := gin.Default()

	// Open routes
	initializeRoutes(router)

	// Linsten routes in 8080 :port
	router.Run(":8080")
}
