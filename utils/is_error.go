package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerErr(err error, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed",
		})
		return
	}
}
