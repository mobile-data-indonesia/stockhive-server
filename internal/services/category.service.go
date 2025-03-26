package services

import (
	"stockhive-server/internal/models"
	"stockhive-server/internal/repositories"
)

type CategoryService interface {
	GetAllCategorys() ([]models.Category, error)
	GetCategoryByID(id string) (models.Category, error)
	CreateCategory(Category *models.Category) error
	UpdateCategory(Category *models.Category) error
	DeleteCategory(Category *models.Category) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) GetAllCategorys() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetCategoryByID(id string) (models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) CreateCategory(Category *models.Category) error {
	return s.repo.Create(Category)
}

func (s *categoryService) UpdateCategory(Category *models.Category) error {
	return s.repo.Update(Category)
}

func (s *categoryService) DeleteCategory(Category *models.Category) error {
	return s.repo.Delete(Category)
}
