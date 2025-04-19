package routes

import (
	"stockhive-server/internal/controllers"
	"stockhive-server/internal/repositories"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

func ActivityLogRoute(r *gin.Engine) {
	activityLogRepo := repositories.NewActivityLogRepository()
	activityLogService := services.NewActivityLogService(activityLogRepo)
	activityLogController := controllers.NewActivityLogController(activityLogService)

	activityLogRoutes := r.Group("/activity-logs")
	{
		activityLogRoutes.GET("/", activityLogController.GetAll)
		activityLogRoutes.GET("/:id", activityLogController.GetByID)
		activityLogRoutes.POST("/", activityLogController.Create)
	}
}
