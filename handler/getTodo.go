package handler

import (
	"errors"
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get a todo based on id
func GetTodoHandler(ctx *gin.Context) {

	// Get query param "id"
	id := ctx.Query("id")

	// Assert at id is given on query
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errMandatoryField("id", "string").Error())
		return
	}

	// Creates a todo based in schema
	todo := schemas.Todo{}

	// Get todo in db according given id
	if err := db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusNotFound, "error todo not found")
			return
		}
		sendError(ctx, http.StatusInternalServerError, "error creating todo")
		return
	}

	// Returns success to client
	sendSuccess(ctx, "create-todo", todo)
}
