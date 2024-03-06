package router

import (
	"github.com/Luis-Honorato/todo_go_api/handler"
	"github.com/gin-gonic/gin"
)

// Declare and Open routes
func initializeRoutes(router *gin.Engine) {
	// Initialize Handlers
	handler.InitializeHandler()

	// Group routes by version [v1], with prefix "/api/v1s"
	v1 := router.Group("/api/v1")
	{
		v1.GET("/todo", handler.GetTodoHandler)
		v1.GET("/todos", handler.ListTodosHandler)
		v1.POST("/todo", handler.CreateTodoHandler)
		v1.PUT("/todo", handler.UpdatetodoHandler)
		v1.DELETE("/todo", handler.DeleteTodo)
	}

}
