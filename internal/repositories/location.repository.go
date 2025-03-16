package repositories

import (
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"
)

type LocationRepository interface {
	FindAll() ([]models.Location, error)
	FindByID(id string) (models.Location, error)
	Create(location *models.Location) error
	Update(location *models.Location) error
	Delete(location *models.Location) error
}

type locationRepository struct{}

func NewLocationRepository() LocationRepository {
	return &locationRepository{}
}

func (r *locationRepository) FindAll() ([]models.Location, error) {
	var locations []models.Location
	err := config.DB.Find(&locations).Error
	return locations, err
}

func (r *locationRepository) FindByID(id string) (models.Location, error) {
	var location models.Location
	err := config.DB.Where("location_id = ?", id).First(&location).Error
	return location, err
}

func (r *locationRepository) Create(location *models.Location) error {
	return config.DB.Create(location).Error
}

func (r *locationRepository) Update(location *models.Location) error {
	return config.DB.Save(location).Error
}

func (r *locationRepository) Delete(location *models.Location) error {
	return config.DB.Delete(location).Error
}
