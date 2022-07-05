package usecase

import (
	"github.com/itsapep/golang-api-with-gin/model"
	"github.com/itsapep/golang-api-with-gin/repository"
)

// type CategoryUsecase interface {
// 	GetCategoryList(ctx *gin.Context)
// 	Register(newCategory *model.Category) error
// }

// type categoryUsecase struct {
// 	categoryRepo repository.CategoryRepository
// }

// // Register implements CategoryUsecase
// func (c *categoryUsecase) Register(newCategory *model.Category) error {
// 	return c.categoryRepo.Add(newCategory)
// }

// // GetCategoryList implements CategoryUsecase
// func (c *categoryUsecase) GetCategoryList(ctx *gin.Context) {
// 	categoryList := c.categoryRepo.GetAll()
// 	utils.Response(http.StatusOK, "OK", categoryList, ctx)
// }

type CategoryUsecase interface {
	Register(newCategory *model.Category) error
	RegisterBulk(newCategory *[]model.Category) error
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

// RegisterBulk implements CategoryUsecase
func (c *categoryUsecase) RegisterBulk(newCategory *[]model.Category) error {
	return c.repo.AddBulk(newCategory)
}

func (c *categoryUsecase) Register(newCategory *model.Category) error {
	return c.repo.Add(newCategory)
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		repo: repo,
	}
}
