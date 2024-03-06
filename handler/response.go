package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Send a error message to client
func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

// Send a success message to client
func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s succesfull", op),
		"data":    data,
	})
}
