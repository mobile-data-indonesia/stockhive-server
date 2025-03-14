package controllers

import (
	"net/http"
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllItems(c *gin.Context) {
	var items []models.Item

	if err := config.DB.Preload("ItemLocation").Preload("Holder").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func GetItemByID(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := config.DB.Preload("ItemLocation").Preload("Holder").First(&item, "item_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, item)
}

func CreateItem(c *gin.Context) {
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if item.ItemLocationID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item_location_id is required"})
		return
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := config.DB.First(&item, "item_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ItemLocationID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item_location_id is required"})
		return
	}

	if err := config.DB.Model(&item).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Item{}, "item_id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
