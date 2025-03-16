package repositories

import (
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"
)

type ItemRepository interface {
	FindAll() ([]models.Item, error)
	FindByID(id string) (models.Item, error)
	Create(item *models.Item) error
	Update(item *models.Item, updated models.Item) error
	Delete(id string) error
}

type itemRepository struct{}

func NewItemRepository() ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) FindAll() ([]models.Item, error) {
	var items []models.Item
	err := config.DB.Preload("ItemLocation").Preload("Holder").Find(&items).Error
	return items, err
}

func (r *itemRepository) FindByID(id string) (models.Item, error) {
	var item models.Item
	err := config.DB.Preload("ItemLocation").Preload("Holder").First(&item, "item_id = ?", id).Error
	return item, err
}

func (r *itemRepository) Create(item *models.Item) error {
	return config.DB.Create(item).Error
}

func (r *itemRepository) Update(item *models.Item, updated models.Item) error {
	return config.DB.Model(item).Updates(updated).Error
}

func (r *itemRepository) Delete(id string) error {
	return config.DB.Delete(&models.Item{}, "item_id = ?", id).Error
}
