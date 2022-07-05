package repository

import "github.com/itsapep/golang-api-with-gin/model"

// type CategoryRepository interface {
// 	Add(newCategory *model.Category) error
// 	GetAll() []model.Category
// 	Update()
// 	Delete()
// }

type CategoryRepository interface {
	Add(newCategory *model.Category) error
	AddBulk(newCategory *[]model.Category) error
	Retrieve() []model.Category
}

type categoryRepository struct {
	categoryDb []model.Category
}

// AddBulk implements CategoryRepository
func (c *categoryRepository) AddBulk(newCategory *[]model.Category) error {
	c.categoryDb = append(c.categoryDb, *newCategory...)
	return nil
}

func (c *categoryRepository) Add(newCategory *model.Category) error {
	c.categoryDb = append(c.categoryDb, *newCategory)
	return nil
}

func (c *categoryRepository) Retrieve() []model.Category {
	return c.categoryDb
}

func NewCategoryRepository() CategoryRepository {
	repo := new(categoryRepository)
	return repo
}
