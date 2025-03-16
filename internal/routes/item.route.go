package routes

import (
	"stockhive-server/internal/controllers"
	"stockhive-server/internal/repositories"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

func ItemRoute(r *gin.Engine) {
	itemRepo := repositories.NewItemRepository()
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)

	itemRoutes := r.Group("/items")
	{
		itemRoutes.GET("/", itemController.GetAll)
		itemRoutes.GET("/:id", itemController.GetByID)
		itemRoutes.POST("/", itemController.Create)
		itemRoutes.PUT("/:id", itemController.Update)
		itemRoutes.DELETE("/:id", itemController.Delete)
	}
}
