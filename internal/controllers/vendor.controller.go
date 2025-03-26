package controllers

import (
	"net/http"
	"stockhive-server/internal/models"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

type VendorController struct {
	service services.VendorService
}

func NewVendorController(service services.VendorService) *VendorController {
	return &VendorController{service}
}

func (ctl *VendorController) GetAll(c *gin.Context) {
	Vendors, err := ctl.service.GetAllVendors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Vendors"})
		return
	}
	c.JSON(http.StatusOK, Vendors)
}

func (ctl *VendorController) GetByID(c *gin.Context) {
	id := c.Param("id")
	Vendor, err := ctl.service.GetVendorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vendor not found"})
		return
	}
	c.JSON(http.StatusOK, Vendor)
}

func (ctl *VendorController) Create(c *gin.Context) {
	var Vendor models.Vendor
	if err := c.ShouldBindJSON(&Vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := ctl.service.CreateVendor(&Vendor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Vendor"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Vendor created successfully"})
}

func (ctl *VendorController) Update(c *gin.Context) {
	id := c.Param("id")
	Vendor, err := ctl.service.GetVendorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vendor not found"})
		return
	}

	if err := c.ShouldBindJSON(&Vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := ctl.service.UpdateVendor(&Vendor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Vendor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vendor updated successfully"})
}

func (ctl *VendorController) Delete(c *gin.Context) {
	id := c.Param("id")
	Vendor, err := ctl.service.GetVendorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vendor not found"})
		return
	}

	if err := ctl.service.DeleteVendor(&Vendor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Vendor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vendor deleted successfully"})
}


