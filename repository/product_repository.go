package repository

import "github.com/itsapep/golang-api-with-gin/model"

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() ([]model.Product, error)
}

type productRepository struct {
	productDb *[]model.Product
}

// Retrieve implements ProductRepository
func (p *productRepository) Retrieve() ([]model.Product, error) {
	return *p.productDb, nil
}

func (p *productRepository) Add(newProduct *model.Product) error {
	*p.productDb = append(*p.productDb, *newProduct)
	return nil
}

func NewProductRepository(productDb *[]model.Product) ProductRepository {
	repo := new(productRepository)
	repo.productDb = productDb
	return repo
}
