package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/usecase"
)

type CategoryController struct {
	router  *gin.Engine
	usecase usecase.CategoryUsecase
}

func (cc *CategoryController) registerNewCategory(ctx *gin.Context) {
	var newCategory model.Category
	_ = ctx.ShouldBindJSON(&newCategory)
	err := cc.usecase.Register(&newCategory)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when creating category",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": newCategory,
	})
}

func (cc *CategoryController) registerNewCategories(ctx *gin.Context) {
	var newCategory []model.Category
	_ = ctx.ShouldBindJSON(&newCategory)
	err := cc.usecase.RegisterBulk(&newCategory)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when creating category",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": newCategory,
	})
}

func NewCategoryController(r *gin.Engine, usecase usecase.CategoryUsecase) *CategoryController {
	controller := CategoryController{
		router:  r,
		usecase: usecase,
	}
	r.POST("/category", controller.registerNewCategory)

	r.POST("/category", controller.registerNewCategories)
	return &controller
}
