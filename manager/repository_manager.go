package manager

import (
	"fmt"

	"github.com/itsapep/golang-api-with-gin/repository"
)

type RepositoryManager interface {
	// collection of all repo
	ProductRepo() repository.ProductRepository
	FileRepo() repository.FileRepository
}

type repositoryManager struct {
	infra Infra
}

// FileRepo implements RepositoryManager
func (r *repositoryManager) FileRepo() repository.FileRepository {
	fmt.Println("filepath infra:", r.infra.FilePath())
	return repository.NewFileRepository(r.infra.FilePath())
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
