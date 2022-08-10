package usecase

import (
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/repository"
)

type FindProductsUsecase interface {
	FindProduct(id string) (model.Product, error)
}

type findProductsUsecase struct {
	repo repository.ProductRepository
}

// GetProducts implements GetProductsUsecase
func (c *findProductsUsecase) FindProduct(id string) (model.Product, error) {
	return c.repo.FindById(id)
}

func NewFindProductsUsecase(repo repository.ProductRepository) FindProductsUsecase {
	return &findProductsUsecase{
		repo: repo,
	}
}
