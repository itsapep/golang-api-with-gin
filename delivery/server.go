package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-api-with-gin/config"
	"github.com/itsapep/golang-api-with-gin/delivery/controller"
	"github.com/itsapep/golang-api-with-gin/repository"
	"github.com/itsapep/golang-api-with-gin/usecase"
)

type appServer struct {
	crProdUc usecase.CreateProductUsecase
	liProdUc usecase.ListProductsUsecase
	engine   *gin.Engine
	host     string
}

func Server() *appServer {
	r := gin.Default()
	productRepo := repository.NewProductRepository()
	crProdUc := usecase.NewCreateProductUsecase(productRepo)
	liProdUc := usecase.NewListProductsUsecase(productRepo)
	c := config.NewConfig()
	host := c.URL
	return &appServer{
		crProdUc: crProdUc,
		liProdUc: liProdUc,
		engine:   r,
		host:     host,
	}
}

// accomodate all controler
func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.crProdUc, a.liProdUc)
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
