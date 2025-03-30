package repositories

import (
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"
)

type AuditLogRepository interface {
	FindAll() ([]models.AuditLog, error)
	FindByID(id string) (models.AuditLog, error)
	Create(auditLog *models.AuditLog) error
}

type auditLogRepository struct{}

func NewAuditLogRepository() AuditLogRepository {
	return &auditLogRepository{}
}

func (r *auditLogRepository) FindAll() ([]models.AuditLog, error) {
	var auditLogs []models.AuditLog
	err := config.DB.Find(&auditLogs).Error
	return auditLogs, err
}

func (r *auditLogRepository) FindByID(id string) (models.AuditLog, error) {
	var auditLog models.AuditLog
	err := config.DB.Where("audit_id = ?", id).First(&auditLog).Error
	return auditLog, err
}

func (r *auditLogRepository) Create(auditLog *models.AuditLog) error {
	return config.DB.Create(auditLog).Error
}
