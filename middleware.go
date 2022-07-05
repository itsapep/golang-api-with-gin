package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/api/auth/login" {
			ctx.Next()
		} else {
			h := AuthHeader{}
			if err := ctx.ShouldBindHeader(&h); err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				ctx.Abort()
			}
			if h.AuthorisationHeader == "ini_token" {
				ctx.Next()
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "invalid token",
				})
				ctx.Abort()
			}
		}
	}
}
