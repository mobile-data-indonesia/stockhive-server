package routes

import (
	"stockhive-server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoute(r *gin.Engine){
	r.POST("/item", controllers.CreateItem)
	r.GET("/item", controllers.GetAllItems)
	r.GET("/item/:id", controllers.GetItemByID)
	r.PUT("/item/:id", controllers.UpdateItem)
	r.DELETE("/item/:id", controllers.DeleteItem)
}
