package handler

import (
	"net/http"

	"github.com/Luis-Honorato/todo_go_api/schemas"
	"github.com/gin-gonic/gin"
)

func ListTodosHandler(ctx *gin.Context) {
	todos := []schemas.Todo{}

	if err := db.Find(&todos).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error list todos")
		return
	}
	sendSuccess(ctx, "list todos", todos)
}
