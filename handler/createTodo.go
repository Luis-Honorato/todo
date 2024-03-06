package handler

import (
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
)

// Create a Todo
func CreateTodoHandler(ctx *gin.Context) {
	// Declare the expected and possible request
	request := CreateTodoRequest{}

	// Read and transform into json received request
	ctx.BindJSON(&request)

	// Validate request
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Create a todo based on request
	todo := schemas.Todo{
		Title:       request.Title,
		Description: request.Description,
		Finished:    *request.Finished,
	}

	// Create a record in DB with given attributes in request
	if err := db.Create(&todo).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error creating todo")
		return
	}

	// Returns success to client
	sendSuccess(ctx, "creating todo", todo)
}
