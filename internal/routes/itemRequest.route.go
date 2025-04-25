package routes

import (
	"stockhive-server/internal/controllers"
	"stockhive-server/internal/repositories"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

func ItemRequestRoute(r *gin.Engine) {
	itemRequestRepo := repositories.NewItemRequestRepository()
	itemRequestService := services.NewItemRequestService(itemRequestRepo)
	itemRequestController := controllers.NewItemRequestController(itemRequestService)

	itemRequestRoutes := r.Group("/item-requests")
	{
		itemRequestRoutes.GET("/", itemRequestController.GetAll)
		itemRequestRoutes.GET("/:id", itemRequestController.GetByID)
		itemRequestRoutes.POST("/", itemRequestController.Create)

	}
}
