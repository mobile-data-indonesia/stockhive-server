package services

import (
	"stockhive-server/internal/models"
	"stockhive-server/internal/repositories"
)

type AuditLogService interface {
	GetAllAuditLogs() ([]models.AuditLog, error)
	GetAuditLogByID(id string) (models.AuditLog, error)
	CreateAuditLog(auditLog *models.AuditLog) error
}

type auditLogService struct {
	repo repositories.AuditLogRepository
}

func NewAuditLogService(repo repositories.AuditLogRepository) AuditLogService {
	return &auditLogService{repo}
}

func (s *auditLogService) GetAllAuditLogs() ([]models.AuditLog, error) {
	return s.repo.FindAll()
}

func (s *auditLogService) GetAuditLogByID(id string) (models.AuditLog, error) {
	return s.repo.FindByID(id)
}

func (s *auditLogService) CreateAuditLog(auditLog *models.AuditLog) error {
	return s.repo.Create(auditLog)
}
