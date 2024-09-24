package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnFailJson(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"msg":     msg,
	})
}

func ReturnSuccessJson(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"msg":     "success",
	})
}
