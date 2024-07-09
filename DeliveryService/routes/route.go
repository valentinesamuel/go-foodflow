package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Delivery Service",
		})
	})

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delivery Service is healthy",
		})
	})

	server.POST("/orders/:id/track", CreateDeliveryTrackingController)
	server.GET("/orders/:id/track", GetDeliveryTrackingStatusController)
}
