package handler

import (
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(ctx *gin.Context) {
	request := CreateTodoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	todo := schemas.Todo{
		Title:       request.Title,
		Description: request.Description,
		Finished:    *request.Finished,
	}

	if err := db.Create(&todo).Error; err != nil {

		sendError(ctx, http.StatusInternalServerError, "error creating todo")
		return
	}

	sendSuccess(ctx, "creating todo", todo)
}
