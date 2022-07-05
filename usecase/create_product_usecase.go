package usecase

import (
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/repository"
)

type CreateProductUsecase interface {
	CreateProduct(newProduct *model.Product) error
}

type createProductUsecase struct {
	repo repository.ProductRepository
}

// CreateProduct implements CreateProductUsecase
func (c *createProductUsecase) CreateProduct(newProduct *model.Product) error {
	return c.repo.Add(newProduct)
}

func NewCreateProductUsecase(repo repository.ProductRepository) CreateProductUsecase {
	return &createProductUsecase{
		repo: repo,
	}
}
