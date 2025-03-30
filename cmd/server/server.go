package server

import (
	"stockhive-server/internal/config"
	"stockhive-server/internal/routes"

	"github.com/gin-gonic/gin"
)


func NewServer() *gin.Engine {
	r := gin.Default()
	config.ConnectDB()

	routes.IndexRoute(r)
	routes.UserRoute(r)
	routes.LocationRoute(r)
	routes.ItemRoute(r)
	routes.VendorRoute(r)
	routes.CategoryRoute(r)
	routes.AuditLogRoute(r)

	return r
}