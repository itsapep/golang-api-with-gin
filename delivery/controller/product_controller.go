package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-api-with-gin/delivery/api"
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/usecase"
	"github.com/itsapep/golang-api-with-gin/utils"
)

type ProductController struct {
	crProdUc usecase.CreateProductUsecase
	liProdUc usecase.ListProductsUsecase
	router   *gin.Engine
	api.BaseAPI
}

func (p *ProductController) createNewProduct(c *gin.Context) {
	var newProduct model.Product
	err := p.ParseRequestBody(c, &newProduct)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.crProdUc.CreateProduct(&newProduct)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newProduct)
}

func (p *ProductController) listProducts(c *gin.Context) {
	products, err := p.liProdUc.ListProducts()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, products)

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
