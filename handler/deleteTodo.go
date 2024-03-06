package handler

import (
	"errors"
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errMandatoryField("id", "int").Error())
		return
	}

	todo := schemas.Todo{}

	if err := db.First(&todo, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "todo not found")
	}

	if err := db.Delete(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusNotFound, "error todo not found")
			return
		}
		sendError(ctx, http.StatusInternalServerError, "error deleting todo")
	}

	sendSuccess(ctx, "delete todo", todo)
}
