package routes

import (
	"stockhive-server/internal/controllers"
	"stockhive-server/internal/repositories"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

func AuditLogRoute(r *gin.Engine) {
	auditLogRepo := repositories.NewAuditLogRepository()
	auditLogService := services.NewAuditLogService(auditLogRepo)
	auditLogController := controllers.NewAuditLogController(auditLogService)

	auditLogRoutes := r.Group("/audit-logs")
	{
		auditLogRoutes.GET("/", auditLogController.GetAll)
		auditLogRoutes.GET("/:id", auditLogController.GetByID)
		auditLogRoutes.POST("/", auditLogController.Create)
	}
}