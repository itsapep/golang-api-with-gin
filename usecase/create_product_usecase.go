package usecase

import (
	"fmt"
	"mime/multipart"

	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/repository"
)

type CreateProductUsecase interface {
	CreateProduct(newProduct *model.Product, file multipart.File, fileExt string) error
}

type createProductUsecase struct {
	repo     repository.ProductRepository
	fileRepo repository.FileRepository
}

// CreateProduct implements CreateProductUsecase
func (c *createProductUsecase) CreateProduct(newProduct *model.Product, file multipart.File, fileExt string) error {
	fileName := fmt.Sprintf("img-%s.%s", newProduct.ProductId, fileExt)
	fileLocation, err := c.fileRepo.Save(file, fileName)
	if err != nil {
		return err
	}
	newProduct.ImgPath = fileLocation
	newProduct.UrlPath = fmt.Sprintf("/product/image/%s", newProduct.ProductId)
	err = c.repo.Add(newProduct)
	if err != nil {
		return err
	}
	return nil
}

func NewCreateProductUsecase(repo repository.ProductRepository, fileRepo repository.FileRepository) CreateProductUsecase {
	return &createProductUsecase{
		repo:     repo,
		fileRepo: fileRepo,
	}
}
