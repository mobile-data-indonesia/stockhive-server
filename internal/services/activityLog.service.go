package services

import (
	"stockhive-server/internal/models"
	"stockhive-server/internal/repositories"
)

type ActivityLogService interface {
	GetAllActivityLogs() ([]models.ActivityLog, error)
	CreateActivityLog(activityLog *models.ActivityLog) error
}

type activityLogService struct {
	repo repositories.ActivityLogRepository
}

func NewActivityLogService(repo repositories.ActivityLogRepository) ActivityLogService {
	return &activityLogService{repo}
}

func (s *activityLogService) GetAllActivityLogs() ([]models.ActivityLog, error) {
	return s.repo.FindAll()
}

func (s *activityLogService) CreateActivityLog(activityLog *models.ActivityLog) error {
	return s.repo.Create(activityLog)
}
