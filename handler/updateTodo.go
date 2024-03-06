package handler

import (
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
)

func UpdatetodoHandler(ctx *gin.Context) {

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errMandatoryField("id", "string").Error())
		return
	}

	request := DeleteTodoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	todo := schemas.Todo{}

	if err := db.First(&todo, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "error todo not found")
		return
	}

	if request.Title != "" {
		todo.Title = request.Title
	}

	if request.Description != "" {
		todo.Description = request.Description
	}
	if request.Finished != nil {
		todo.Finished = *request.Finished
	}

	if err := db.Save(&todo).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error saving todo")
		return
	}

	sendSuccess(ctx, "create-todo", todo)
}
