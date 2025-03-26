package routes

import (
	"stockhive-server/internal/controllers"
	"stockhive-server/internal/repositories"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

func VendorRoute(r *gin.Engine) {
	vendorRepo := repositories.NewVendorRepository()
	vendorService := services.NewVendorService(vendorRepo)
	vendorController := controllers.NewVendorController(vendorService)

	vendorRoutes := r.Group("/vendors")
	{
		vendorRoutes.GET("/", vendorController.GetAll)
		vendorRoutes.GET("/:id", vendorController.GetByID)
		vendorRoutes.POST("/", vendorController.Create)
		vendorRoutes.PUT("/:id",vendorController.Update)
		vendorRoutes.DELETE("/:id", vendorController.Delete)
	}
}
