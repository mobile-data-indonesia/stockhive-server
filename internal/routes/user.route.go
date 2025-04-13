package routes

import (
	"stockhive-server/internal/controllers"
	"stockhive-server/internal/repositories"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	userController := controllers.NewUserController(service)

	// Corrected handlers
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.POST("/refresh", userController.RefreshToken)
	r.POST("/change-password", userController.ChangePassword)
	r.GET("/users", userController.GetAllUsers)
	r.PUT("/users/:id", userController.UpdateUser)
	r.GET("/users/:id", userController.GetUserByID)
	r.DELETE("/users/:id", userController.DeleteUser)
}
