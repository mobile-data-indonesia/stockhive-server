package repositories

import (
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"
)

type ItemRequestRepository interface {
	FindAll() ([]models.ItemRequest, error)
	FindByID(id string) (models.ItemRequest, error)
	Create(itemRequest *models.ItemRequest) error
	FindByRequesterID(requesterID string) ([]models.ItemRequest, error)
	FindByStatus(status string) ([]models.ItemRequest, error)
	FindByRequesterIDAndStatus(requesterID, status string) ([]models.ItemRequest, error)
}

type itemRequestRepository struct{}

func NewItemRequestRepository() ItemRequestRepository {
	return &itemRequestRepository{}
}

func (r *itemRequestRepository) FindAll() ([]models.ItemRequest, error) {
	var itemRequests []models.ItemRequest
	err := config.DB.Preload("Requester").Find(&itemRequests).Error
	return itemRequests, err
}

func (r *itemRequestRepository) FindByID(id string) (models.ItemRequest, error) {
	var itemRequest models.ItemRequest
	err := config.DB.Preload("Requester").Where("item_request_id = ?", id).First(&itemRequest).Error
	return itemRequest, err
}

func (r *itemRequestRepository) Create(itemRequest *models.ItemRequest) error {
	return config.DB.Create(itemRequest).Error
}

func (r *itemRequestRepository) FindByRequesterID(requesterID string) ([]models.ItemRequest, error) {
	var itemRequests []models.ItemRequest
	err := config.DB.Preload("Requester").Where("requester_id = ?", requesterID).Find(&itemRequests).Error
	return itemRequests, err
}

func (r *itemRequestRepository) FindByStatus(status string) ([]models.ItemRequest, error) {
	var itemRequests []models.ItemRequest
	err := config.DB.Preload("Requester").Where("request_status = ?", status).Find(&itemRequests).Error
	return itemRequests, err
}

func (r *itemRequestRepository) FindByRequesterIDAndStatus(requesterID, status string) ([]models.ItemRequest, error) {
	var itemRequests []models.ItemRequest
	err := config.DB.Preload("Requester").Where("requester_id = ? AND request_status = ?", requesterID, status).Find(&itemRequests).Error
	return itemRequests, err
}
