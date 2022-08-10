package controller

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-api-with-gin/delivery/api"
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/usecase"
)

type ProductController struct {
	crProdUc usecase.CreateProductUsecase
	liProdUc usecase.ListProductsUsecase
	fiProdUc usecase.FindProductsUsecase
	router   *gin.Engine
	api.BaseAPI
}

func (p *ProductController) createNewProduct(c *gin.Context) {
	productId := c.PostForm("productId")
	productName := c.PostForm("productName")
	file, fileHeader, err := c.Request.FormFile("productImg")
	if err != nil {
		p.Failed(c, errors.New("failed getting file"))
		return
	}
	log.Println(fileHeader.Filename) // gambar.jpg
	fileName := strings.Split(fileHeader.Filename, ".")
	if len(fileName) != 2 {
		p.Failed(c, errors.New("unrecognised file error"))
		return
	}
	var newProduct = model.Product{
		ProductId:   productId,
		ProductName: productName,
	}

	err = p.crProdUc.CreateProduct(&newProduct, file, fileName[1])
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

func (p *ProductController) findImageById(c *gin.Context) {
	productId := c.Param("id")
	product, err := p.fiProdUc.FindProduct(productId)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.SuccessDownload(c, product.ImgPath)

}

func NewProductController(router *gin.Engine, crProdUc usecase.CreateProductUsecase, liProdUc usecase.ListProductsUsecase, fiProdUc usecase.FindProductsUsecase) *ProductController {
	// here lies all request method we need
	controller := ProductController{
		crProdUc: crProdUc,
		liProdUc: liProdUc,
		fiProdUc: fiProdUc,
		router:   router,
	}
	router.POST("/product", controller.createNewProduct)

	router.GET("/product", controller.listProducts)
	router.GET("/product/image/:id", controller.findImageById)

	return &controller
}
