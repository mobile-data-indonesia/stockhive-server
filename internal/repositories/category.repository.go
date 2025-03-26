package repositories

import (
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	FindByID(id string) (models.Category, error)
	Create(Category *models.Category) error
	Update(Category *models.Category) error
	Delete(Category *models.Category) error
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
	var Categorys []models.Category
	err := config.DB.Find(&Categorys).Error
	return Categorys, err
}

func (r *categoryRepository) FindByID(id string) (models.Category, error) {
	var Category models.Category
	err := config.DB.Where("Category_id = ?", id).First(&Category).Error
	return Category, err
}

func (r *categoryRepository) Create(Category *models.Category) error {
	return config.DB.Create(Category).Error
}

func (r *categoryRepository) Update(Category *models.Category) error {
	return config.DB.Save(Category).Error
}

func (r *categoryRepository) Delete(Category *models.Category) error {
	return config.DB.Delete(Category).Error
}
