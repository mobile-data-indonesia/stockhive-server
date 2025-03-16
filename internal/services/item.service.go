package services

import (
	"errors"
	"stockhive-server/internal/models"
	"stockhive-server/internal/repositories"

	"gorm.io/gorm"
)

type ItemService interface {
	GetAllItems() ([]models.Item, error)
	GetItemByID(id string) (models.Item, error)
	CreateItem(item *models.Item) error
	UpdateItem(id string, updated models.Item) (models.Item, error)
	DeleteItem(id string) error
}

type itemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{repo}
}

func (s *itemService) GetAllItems() ([]models.Item, error) {
	return s.repo.FindAll()
}

func (s *itemService) GetItemByID(id string) (models.Item, error) {
	return s.repo.FindByID(id)
}

func (s *itemService) CreateItem(item *models.Item) error {
	if item.ItemLocationID == 0 {
		return errors.New("item_location_id is required")
	}
	return s.repo.Create(item)
}

func (s *itemService) UpdateItem(id string, updated models.Item) (models.Item, error) {
	if updated.ItemLocationID == 0 {
		return models.Item{}, errors.New("item_location_id is required")
	}

	item, err := s.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Item{}, err
		}
		return models.Item{}, err
	}

	err = s.repo.Update(&item, updated)
	return item, err
}

func (s *itemService) DeleteItem(id string) error {
	return s.repo.Delete(id)
}
