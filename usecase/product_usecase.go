package usecase

import (
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/repository"
)

// type ProductUsecase interface {
// 	PostProduct(ctx *gin.Context)
// 	GetProductById(ctx *gin.Context)
// 	GetProductByCategory(ctx *gin.Context) // return list
// 	// GetProductCategoryById(ctx *gin.Context)
// 	DeleteProduct(ctx *gin.Context) // update is_active
// }

// type productUsecase struct {
// 	productRepo repository.ProductRepository
// }

type ProductUsecase interface {
	Register(newProduct *model.Product) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func (a *productUsecase) Register(newProduct *model.Product) error {
	return a.repo.Add(newProduct)
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}
