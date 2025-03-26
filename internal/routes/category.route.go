package routes

import (
	"stockhive-server/internal/controllers"
	"stockhive-server/internal/repositories"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(r *gin.Engine) {
	categoryRepo := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	categoryRoutes := r.Group("/categories")
	{
		categoryRoutes.GET("/", categoryController.GetAll)
		categoryRoutes.GET("/:id", categoryController.GetByID)
		categoryRoutes.POST("/", categoryController.Create)
		categoryRoutes.PUT("/:id", categoryController.Update)
		categoryRoutes.DELETE("/:id", categoryController.Delete)
	}
}
