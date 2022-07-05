package manager

import "github.com/itsapep/golang-api-with-gin/repository"

type RepositoryManager interface {
	// collection of all repo
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
}

// ProductRepo implements RepositoryManager
func (*repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository()
}

func NewRepositoryManager() RepositoryManager {
	return &repositoryManager{}
}
