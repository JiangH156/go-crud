package repository

import (
	"gorm.io/gorm"
	"john/gin-curd/common"
	"john/gin-curd/models"
)

type ICategoryRepository interface {
	Create(name string) error
	Delete(id int) error
	UpdateById(id int, name string) (*models.Category, error)
	Query(id int) (*models.Category, error)
}

type CategoryRepository struct {
	DB *gorm.DB
}

func NewICategoryRepository() CategoryRepository {
	return CategoryRepository{
		DB: common.GetDB(),
	}
}

func (c CategoryRepository) SelectById(id int) (*models.Category, error) {
	category := models.Category{}
	if err := c.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) Create(name string) error {
	category := models.Category{
		Name: name,
	}

	if err := c.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (c CategoryRepository) Delete(id int) error {
	var category = models.Category{}
	if err := c.DB.Where("id = ?", id).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func (c CategoryRepository) UpdateById(id int, name string) (*models.Category, error) {
	var category = &models.Category{ID: uint(id)}
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c CategoryRepository) Query(id int) (*models.Category, error) {
	var category = models.Category{}
	if err := c.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
