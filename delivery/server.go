package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-api-with-gin/config"
	"github.com/itsapep/golang-api-with-gin/delivery/controller"
	"github.com/itsapep/golang-api-with-gin/manager"
)

type appServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(&appConfig)
	repoManager := manager.NewRepositoryManager(infra)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := appConfig.URL
	return &appServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}

// accomodate all controler
func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.usecaseManager.ProductRegistrationUsecase(), a.usecaseManager.ListProductUsecase(), a.usecaseManager.FindProductsUsecase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
