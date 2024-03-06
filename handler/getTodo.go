package handler

import (
	"errors"
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTodoHandler(ctx *gin.Context) {
	todo := schemas.Todo{}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errMandatoryField("id", "string").Error())
		return
	}

	if err := db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusNotFound, "error todo not found")
			return
		}
		sendError(ctx, http.StatusInternalServerError, "error creating todo")
		return
	}

	sendSuccess(ctx, "create-todo", todo)
}
