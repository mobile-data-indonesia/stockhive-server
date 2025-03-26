package services

import (
	"stockhive-server/internal/models"
	"stockhive-server/internal/repositories"
)

type VendorService interface {
	GetAllVendors() ([]models.Vendor, error)
	GetVendorByID(id string) (models.Vendor, error)
	CreateVendor(Vendor *models.Vendor) error
	UpdateVendor(Vendor *models.Vendor) error
	DeleteVendor(Vendor *models.Vendor) error
}

type vendorService struct {
	repo repositories.VendorRepository
}

func NewVendorService(repo repositories.VendorRepository) VendorService {
	return &vendorService{repo}
}

func (s *vendorService) GetAllVendors() ([]models.Vendor, error) {
	return s.repo.FindAll()
}

func (s *vendorService) GetVendorByID(id string) (models.Vendor, error) {
	return s.repo.FindByID(id)
}

func (s *vendorService) CreateVendor(Vendor *models.Vendor) error {
	return s.repo.Create(Vendor)
}

func (s *vendorService) UpdateVendor(Vendor *models.Vendor) error {
	return s.repo.Update(Vendor)
}

func (s *vendorService) DeleteVendor(Vendor *models.Vendor) error {
	return s.repo.Delete(Vendor)
}
