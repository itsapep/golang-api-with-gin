package utils

import "github.com/gin-gonic/gin"

func Response(status int, string string, data interface{}, ctx *gin.Context) {
	ctx.JSON(status, gin.H{
		"message": string,
		"data":    data,
	})
}
