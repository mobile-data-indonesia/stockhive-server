package routes

import (
	"stockhive-server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func LocationRoute(r *gin.Engine) {
	r.GET("/location", controllers.GetAllLocations)
	r.GET("/location/:id", controllers.GetLocation)
	r.POST("/location", controllers.CreateLocation)
	r.PUT("/location/:id", controllers.UpdateLocation)
	r.DELETE("/location/:id", controllers.DeleteLocation)
}
