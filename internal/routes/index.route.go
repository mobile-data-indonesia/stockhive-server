package routes

import (
	"stockhive-server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func IndexRoute(r *gin.Engine) {
	r.GET("/", controllers.Index)
}
