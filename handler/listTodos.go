package handler

import (
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
)

// List all todos in db
func ListTodosHandler(ctx *gin.Context) {
	// Create a list of todos based in schema
	todos := []schemas.Todo{}

	// Find all todos in db
	if err := db.Find(&todos).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error list todos")
		return
	}

	// Returns succes to client
	sendSuccess(ctx, "list todos", todos)
}
