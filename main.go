package main

import "github.com/itsapep/golang-api-with-gin/delivery"

func main() {
	delivery.Server().Run()
}

// func main() {

// routerEngine := gin.Default()
// // routerEngine.Use(AuthMiddleware())
// prodRepo := repository.NewProductRepository()
// prodUsc := usecase.NewProductUseCase(prodRepo)

// catRepo := repository.NewCategoryRepository()
// catUsc := usecase.NewCategoryUseCase(catRepo)

// routerEngine.POST("/product", prodUsc.PostProduct)

// routerEngine.GET("/category", catUsc.GetCategoryList)

// routerEngine.GET("/", func(ctx *gin.Context) {
// 	ctx.String(200, "Healthy check")
// })

// routerEngine.GET("/greeting/:name", greeting)

// routerEngine.GET("/greeting/:name", greetingAddress)

// routerEngine.POST("/login", login)

// rgAuth := routerEngine.Group("/api/auth")
// rgMaster := routerEngine.Group("/api/master")
// rgMaster.GET("/greeting/:name", greetingAddress)
// rgAuth.POST("/login", login)

// err := routerEngine.Run(":8888") // default port is 8080
// if err != nil {
// 	panic(err)
// }
// }

// func greeting(ctx *gin.Context) {
// 	name := ctx.Param("name") // with parameter :name at path /greeting
// 	ctx.String(http.StatusOK, "Greeting path, hello ...%s", name)
// }

// // query param/string in path->?key=value
// func greetingAddress(ctx *gin.Context) {
// 	name := ctx.Param("name") // with parameter :name at path /greeting
// 	kec := ctx.Query("kecamatan")
// 	kel := ctx.Query("kelurahan")
// 	ctx.String(http.StatusOK, "Greeting path, hello%s \nkamu ada di kecamatan %s kelurahan %s", name, kec, kel)
// }

// func login(ctx *gin.Context) {
// 	var uc UserCredential
// 	// model binding & validation
// 	// ShouldBind and ShouldBindJSON require struct tag binding:required
// 	//
// 	if err := ctx.BindJSON(&uc); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		fmt.Println(err.Error())
// 	} else {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"token": "ini_token",
// 		})
// 	}
// }
