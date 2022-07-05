package manager

import (
	"github.com/itsapep/golang-api-with-gin/repository"
)

type RepositoryManager interface {
	// collection of all repo
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	infra Infra
}

// ProductRepo implements RepositoryManager
func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
