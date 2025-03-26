package repositories

import (
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"
)

type VendorRepository interface {
	FindAll() ([]models.Vendor, error)
	FindByID(id string) (models.Vendor, error)
	Create(Vendor *models.Vendor) error
	Update(Vendor *models.Vendor) error
	Delete(Vendor *models.Vendor) error
}

type vendorRepository struct{}

func NewVendorRepository() VendorRepository {
	return &vendorRepository{}
}

func (r *vendorRepository) FindAll() ([]models.Vendor, error) {
	var vendors []models.Vendor
	err := config.DB.Find(&vendors).Error
	return vendors, err
}

func (r *vendorRepository) FindByID(id string) (models.Vendor, error) {
	var vendor models.Vendor
	err := config.DB.Where("vendor_id = ?", id).First(&vendor).Error
	return vendor, err
}

func (r *vendorRepository) Create(vendor *models.Vendor) error {
	return config.DB.Create(vendor).Error
}

func (r *vendorRepository) Update(vendor *models.Vendor) error {
	return config.DB.Save(vendor).Error
}

func (r *vendorRepository) Delete(vendor *models.Vendor) error {
	return config.DB.Delete(vendor).Error
}
