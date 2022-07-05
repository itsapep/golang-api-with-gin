package manager

import (
	"github.com/itsapep/golang-api-with-gin/repository"
)

type RepositoryManager interface {
	// collection of all repo
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	productDb repository.ProductRepository
}

// ProductRepo implements RepositoryManager
func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return r.productDb
}

func NewRepositoryManager() RepositoryManager {
	productDb := repository.NewProductRepository()
	return &repositoryManager{
		productDb: productDb,
	}
}
