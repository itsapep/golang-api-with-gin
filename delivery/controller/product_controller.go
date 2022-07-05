package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/usecase"
)

type ProductController struct {
	crProdUc usecase.CreateProductUsecase
	liProdUc usecase.ListProductsUsecase
	router   *gin.Engine
}

func (p *ProductController) createNewProduct(c *gin.Context) {
	var newProduct *model.Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		err := p.crProdUc.CreateProduct(newProduct)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating Product",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": newProduct,
		})
	}
}

func (p *ProductController) listProducts(c *gin.Context) {
	products, err := p.liProdUc.ListProducts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "FAILED",
			"message": "Error when retrieve Product",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": products,
	})

}

func NewProductController(router *gin.Engine, crProdUc usecase.CreateProductUsecase, liProdUc usecase.ListProductsUsecase) *ProductController {
	// here lies all request method we need
	controller := ProductController{
		crProdUc: crProdUc,
		liProdUc: liProdUc,
		router:   router,
	}
	router.POST("/product", controller.createNewProduct)

	router.GET("/product", controller.listProducts)

	return &controller
}
