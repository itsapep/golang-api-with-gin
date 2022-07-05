package usecase

import (
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/repository"
)

type ListProductsUsecase interface {
	ListProducts() ([]model.Product, error)
}

type listProductsUsecase struct {
	repo repository.ProductRepository
}

// GetProducts implements GetProductsUsecase
func (c *listProductsUsecase) ListProducts() ([]model.Product, error) {
	return c.repo.Retrieve()
}

func NewListProductsUsecase(repo repository.ProductRepository) ListProductsUsecase {
	return &listProductsUsecase{
		repo: repo,
	}
}
