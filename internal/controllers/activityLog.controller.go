package controllers

import (
	"fmt"
	"net/http"
	"stockhive-server/internal/models"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

type ActivityLogController struct {
	service services.ActivityLogService
}

func NewActivityLogController(service services.ActivityLogService) *ActivityLogController {
	return &ActivityLogController{service}
}

func (ctl *ActivityLogController) GetAll(c *gin.Context) {
	activityLogs, err := ctl.service.GetAllActivityLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, activityLogs)
}

func (ctl *ActivityLogController) GetByID(c *gin.Context) {
	id := c.Param("id")
	activityLog, err := ctl.service.GetActivityLogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity log not found"})
		return
	}
	c.JSON(http.StatusOK, activityLog)
}

func (ctl *ActivityLogController) Create(c *gin.Context) {
	var activityLog models.ActivityLog
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&activityLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctl.service.CreateActivityLog(&activityLog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Activity log created successfully"})
}
