package services

import (
	"stockhive-server/internal/models"
	"stockhive-server/internal/repositories"
)

type ItemRequestService interface {
	GetAllItemRequests() ([]models.ItemRequest, error)
	GetItemRequestByID(id string) (models.ItemRequest, error)
	CreateItemRequest(itemRequest *models.ItemRequest) error
	GetItemRequestsByRequesterID(requesterID string) ([]models.ItemRequest, error)
	GetItemRequestsByStatus(status string) ([]models.ItemRequest, error)
	GetItemRequestsByRequesterIDAndStatus(requesterID, status string) ([]models.ItemRequest, error)
}

type itemRequestService struct {
	repo repositories.ItemRequestRepository
}

func NewItemRequestService(repo repositories.ItemRequestRepository) ItemRequestService {
	return &itemRequestService{repo}
}

func (s *itemRequestService) GetAllItemRequests() ([]models.ItemRequest, error) {
	return s.repo.FindAll()
}

func (s *itemRequestService) GetItemRequestByID(id string) (models.ItemRequest, error) {
	return s.repo.FindByID(id)
}

func (s *itemRequestService) CreateItemRequest(itemRequest *models.ItemRequest) error {
	return s.repo.Create(itemRequest)
}

func (s *itemRequestService) GetItemRequestsByRequesterID(requesterID string) ([]models.ItemRequest, error) {
	return s.repo.FindByRequesterID(requesterID)
}

func (s *itemRequestService) GetItemRequestsByStatus(status string) ([]models.ItemRequest, error) {
	return s.repo.FindByStatus(status)
}

func (s *itemRequestService) GetItemRequestsByRequesterIDAndStatus(requesterID, status string) ([]models.ItemRequest, error) {
	return s.repo.FindByRequesterIDAndStatus(requesterID, status)
}
