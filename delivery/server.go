package delivery

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")
	host := fmt.Sprintf("%s:%s", apiHost, apiPort)
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
