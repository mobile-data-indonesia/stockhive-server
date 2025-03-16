package services

import (
	"stockhive-server/internal/models"
	"stockhive-server/internal/repositories"
)

type LocationService interface {
	GetAllLocations() ([]models.Location, error)
	GetLocationByID(id string) (models.Location, error)
	CreateLocation(location *models.Location) error
	UpdateLocation(location *models.Location) error
	DeleteLocation(location *models.Location) error
}

type locationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) LocationService {
	return &locationService{repo}
}

func (s *locationService) GetAllLocations() ([]models.Location, error) {
	return s.repo.FindAll()
}

func (s *locationService) GetLocationByID(id string) (models.Location, error) {
	return s.repo.FindByID(id)
}

func (s *locationService) CreateLocation(location *models.Location) error {
	return s.repo.Create(location)
}

func (s *locationService) UpdateLocation(location *models.Location) error {
	return s.repo.Update(location)
}

func (s *locationService) DeleteLocation(location *models.Location) error {
	return s.repo.Delete(location)
}
