package controllers

import (
	"net/http"
	"stockhive-server/internal/models"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemRequestController struct {
	service services.ItemRequestService
}

func NewItemRequestController(service services.ItemRequestService) *ItemRequestController {
	return &ItemRequestController{service}
}

func (ctl *ItemRequestController) GetAll(c *gin.Context) {
	itemRequests, err := ctl.service.GetAllItemRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item requests"})
		return
	}
	c.JSON(http.StatusOK, itemRequests)
}

func (ctl *ItemRequestController) GetByID(c *gin.Context) {
	id := c.Param("id")
	itemRequest, err := ctl.service.GetItemRequestByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item request not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, itemRequest)
}

func (ctl *ItemRequestController) Create(c *gin.Context) {
	var itemRequest models.ItemRequest
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctl.service.CreateItemRequest(&itemRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, itemRequest)
}

func (ctl *ItemRequestController) GetByRequesterID(c *gin.Context) {
	requesterID := c.Param("requester_id")
	itemRequests, err := ctl.service.GetItemRequestsByRequesterID(requesterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item requests"})
		return
	}
	c.JSON(http.StatusOK, itemRequests)
}

func (ctl *ItemRequestController) GetByStatus(c *gin.Context) {
	status := c.Param("status")
	itemRequests, err := ctl.service.GetItemRequestsByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item requests"})
		return
	}
	c.JSON(http.StatusOK, itemRequests)
}

func (ctl *ItemRequestController) GetByRequesterIDAndStatus(c *gin.Context) {
	requesterID := c.Param("requester_id")
	status := c.Param("status")
	itemRequests, err := ctl.service.GetItemRequestsByRequesterIDAndStatus(requesterID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item requests"})
		return
	}
	c.JSON(http.StatusOK, itemRequests)
}
