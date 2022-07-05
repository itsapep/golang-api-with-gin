package manager

import "github.com/itsapep/golang-api-with-gin/usecase"

type UsecaseManager interface {
	ProductRegistrationUsecase() usecase.CreateProductUsecase
	ListProductUsecase() usecase.ListProductsUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

// ListProductUsecase implements UsecaseManager
func (u *usecaseManager) ListProductUsecase() usecase.ListProductsUsecase {
	return usecase.NewListProductsUsecase(u.repoManager.ProductRepo())
}

// ProductRegistrationUsecase implements UsecaseManager
func (u *usecaseManager) ProductRegistrationUsecase() usecase.CreateProductUsecase {
	return usecase.NewCreateProductUsecase(u.repoManager.ProductRepo())
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
