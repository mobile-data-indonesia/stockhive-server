package controllers

import (
	"fmt"
	"net/http"
	"stockhive-server/internal/models"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

type AuditLogController struct {
	service services.AuditLogService
}

func NewAuditLogController(service services.AuditLogService) *AuditLogController {
	return &AuditLogController{service}
}

func (ctl *AuditLogController) GetAll(c *gin.Context) {
	auditLogs, err := ctl.service.GetAllAuditLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch audit logs"})
		return
	}
	c.JSON(http.StatusOK, auditLogs)
}

func (ctl *AuditLogController) GetByID(c *gin.Context) {
	id := c.Param("id")
	auditLog, err := ctl.service.GetAuditLogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Audit log not found"})
		return
	}
	c.JSON(http.StatusOK, auditLog)
}

func (ctl *AuditLogController) Create(c *gin.Context) {
	var auditLog models.AuditLog
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&auditLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctl.service.CreateAuditLog(&auditLog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create audit log"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Audit log created successfully"})
}
