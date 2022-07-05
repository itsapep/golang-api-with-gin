package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/usecase"
)

type ProductController struct {
	router  *gin.Engine
	usecase usecase.ProductUsecase
}

func (cc *ProductController) registerNewProduct(ctx *gin.Context) {
	productId := ctx.PostForm("productId")
	productName := ctx.PostForm("productName")
	newProduct := model.Product{
		ProductId:   productId,
		ProductName: productName,
		Category:    model.Category{},
		IsActive:    false,
	}
	err := cc.usecase.Register(&newProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when creating product",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": newProduct,
	})
}

func NewProductController(r *gin.Engine, usecase usecase.ProductUsecase) *ProductController {
	controller := ProductController{
		router:  r,
		usecase: usecase,
	}
	r.POST("/product", controller.registerNewProduct)
	return &controller
}
