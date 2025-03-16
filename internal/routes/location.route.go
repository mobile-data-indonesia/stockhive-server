package routes

import (
	"stockhive-server/internal/controllers"
	"stockhive-server/internal/repositories"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

func LocationRoute(r *gin.Engine) {
	locationRepo := repositories.NewLocationRepository()
	locationService := services.NewLocationService(locationRepo)
	locationController := controllers.NewLocationController(locationService)

	locationRoutes := r.Group("/locations")
	{
		locationRoutes.GET("/", locationController.GetAll)
		locationRoutes.GET("/:id", locationController.GetByID)
		locationRoutes.POST("/", locationController.Create)
		locationRoutes.PUT("/:id", locationController.Update)
		locationRoutes.DELETE("/:id", locationController.Delete)
	}
}
