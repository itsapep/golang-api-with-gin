package manager

import (
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/repository"
)

type RepositoryManager interface {
	// collection of all repo
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	productDb []model.Product
}

// ProductRepo implements RepositoryManager
func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(&r.productDb)
}

func NewRepositoryManager() RepositoryManager {
	return &repositoryManager{}
}
