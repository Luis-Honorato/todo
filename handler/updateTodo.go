package handler

import (
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
)

// Update a Todo according given id
func UpdatetodoHandler(ctx *gin.Context) {

	// Get query param id
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errMandatoryField("id", "string").Error())
		return
	}

	// Declare the expected and possible request
	request := DeleteTodoRequest{}

	// Read and transform into json received request
	ctx.BindJSON(&request)

	// Validate request
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Create a Todo based in schema
	todo := schemas.Todo{}

	// Get first occurency of Todo in db based with id
	if err := db.First(&todo, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "error todo not found")
		return
	}

	// Update todo if attribute is not empty in request

	if request.Title != "" {
		todo.Title = request.Title
	}

	if request.Description != "" {
		todo.Description = request.Description
	}
	if request.Finished != nil {
		todo.Finished = *request.Finished
	}

	// Save in db updated todo
	if err := db.Save(&todo).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error saving todo")
		return
	}

	// Returns success to client
	sendSuccess(ctx, "create-todo", todo)
}
