package handler

import (
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
)

// Delete a todo based on id
func DeleteTodo(ctx *gin.Context) {
	// Get querry param "id"
	id := ctx.Query("id")

	// Assert at id is given on query
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errMandatoryField("id", "int").Error())
		return
	}

	// Create an todo based on schema
	todo := schemas.Todo{}

	// Get todo in db according given id
	if err := db.First(&todo, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "todo not found")
		return
	}

	// Deletes todo in db according given id
	if err := db.Delete(&todo, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error deleting todo")
		return
	}

	// Returns success to client
	sendSuccess(ctx, "delete todo", todo)
}
