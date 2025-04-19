package repositories

import (
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"
)

type ActivityLogRepository interface {
	FindAll() ([]models.ActivityLog, error)
	FindByID(id string) (models.ActivityLog, error)
	Create(activityLog *models.ActivityLog) error
}

type activityLogRepository struct{}

func NewActivityLogRepository() ActivityLogRepository {
	return &activityLogRepository{}
}

func (r *activityLogRepository) FindAll() ([]models.ActivityLog, error) {
	var activityLogs []models.ActivityLog
	err := config.DB.Preload("Actor").Find(&activityLogs).Error
	return activityLogs, err
}

func (r *activityLogRepository) FindByID(id string) (models.ActivityLog, error) {
	var activityLog models.ActivityLog
	err := config.DB.Preload("Actor").Where("activity_id = ?", id).First(&activityLog).Error
	return activityLog, err
}

func (r *activityLogRepository) Create(activityLog *models.ActivityLog) error {
	return config.DB.Create(activityLog).Error
}
