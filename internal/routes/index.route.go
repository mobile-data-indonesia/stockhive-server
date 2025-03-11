package routes

import (
	"stockhive-server/internal/controllers"
	"stockhive-server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func IndexRoute(r *gin.Engine) {
	r.GET("/", middlewares.JWTMiddleware("access"),middlewares.RoleMiddleware("admin"),controllers.Index)
}
