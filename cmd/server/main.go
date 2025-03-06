package server

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ponggggggggggggggs",
		})
	})
	return r
}