package controllers

import (
	"net/http"
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
	var locations []models.Location

	if err := config.DB.Find(&locations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch locations"})
		return
	}

	c.JSON(http.StatusOK, locations)
}

func GetLocation(c *gin.Context) {
	var location models.Location
	// fmt.Println(c.Param("id"))
	id := c.Param("id")
	// fmt.Println(id)
	if err := config.DB.Where("location_id = ?", id).First(&location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, location)
}

func CreateLocation(c *gin.Context) {
	var location models.Location

	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := config.DB.Create(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create location"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Location created successfully",
	})
}

func UpdateLocation(c *gin.Context) {
	var location models.Location
	id := c.Param("id")

	if err := config.DB.Where("location_id = ?", id).First(&location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	config.DB.Save(&location)

	c.JSON(http.StatusOK, gin.H{
		"message": "Location updated successfully",
	})
}

func DeleteLocation(c *gin.Context) {
	var location models.Location
	id := c.Param("id")

	if err := config.DB.Where("location_id = ?", id).First(&location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	config.DB.Delete(&location)

	c.JSON(http.StatusOK, gin.H{
		"message": "Location deleted successfully",
	})
}