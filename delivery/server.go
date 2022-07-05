package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-api-with-gin/delivery/controller"
	"github.com/itsapep/golang-api-with-gin/repository"
	"github.com/itsapep/golang-api-with-gin/usecase"
)

type appServer struct {
	productUsecase  usecase.ProductUsecase
	categoryUsecase usecase.CategoryUsecase
	engine          *gin.Engine
	host            string
}

func (p *appServer) initHandlers() {
	controller.NewProductController(p.engine, p.productUsecase)
	controller.NewCategoryController(p.engine, p.categoryUsecase)
}

func (p *appServer) Run() {
	p.initHandlers()
	err := p.engine.Run(p.host)
	if err != nil {
		panic(err)
	}
}

func Server() *appServer {
	r := gin.Default()
	prodRepo := repository.NewProductRepository()
	catRepo := repository.NewCategoryRepository()
	prodUsecase := usecase.NewProductUsecase(prodRepo)
	catUsecase := usecase.NewCategoryUsecase(catRepo)
	host := fmt.Sprintf("%s:%s", "localhost", "8888")
	return &appServer{
		productUsecase:  prodUsecase,
		categoryUsecase: catUsecase,
		engine:          r,
		host:            host,
	}
}
