package services

import (
	"stockhive-server/internal/models"
	"stockhive-server/internal/repositories"
)

type ActivityLogService interface {
	GetAllActivityLogs() ([]models.ActivityLog, error)
	GetActivityLogByID(id string) (models.ActivityLog, error)
	CreateActivityLog(activityLog *models.ActivityLog) error
}

func NewActivityLogService(repo repositories.ActivityLogRepository) ActivityLogService {
	return &activityLogService{repo}
}

type activityLogService struct {
	repo repositories.ActivityLogRepository
}

func (s *activityLogService) GetAllActivityLogs() ([]models.ActivityLog, error) {
	activityLogs, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return activityLogs, nil
}

func (s *activityLogService) GetActivityLogByID(id string) (models.ActivityLog, error) {
	activityLog, err := s.repo.FindByID(id)
	if err != nil {
		return models.ActivityLog{}, err
	}
	return activityLog, nil
}

func (s *activityLogService) CreateActivityLog(activityLog *models.ActivityLog) error {
	err := s.repo.Create(activityLog)
	if err != nil {
		return err
	}
	return nil
}