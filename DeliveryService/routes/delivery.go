package routes

import (
	"DeliveryService/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valentinesamuel/go-foodflow/utilities/models"
)

func CreateDeliveryTrackingController(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	deliveryTracking := models.DeliveryTracking{
		OrderID:                 uint(orderID),
		CurrentStatus:           "Pending",
		EstimatedTimeOfDelivery: time.Date(2024, time.January, 10, 23, 0, 0, 0, time.UTC),
	}
	createdDeliveryTracking, err := service.CreateDeliveryTrackingService(&deliveryTracking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create delivery tracking"})
		return
	}
	c.JSON(http.StatusOK, createdDeliveryTracking)
}

func GetDeliveryTrackingStatusController(c *gin.Context) {

	deliveryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid delivery ID"})
		return
	}

	deliveryTracking, err := service.GetDeliveryTrackingStatusService(uint(deliveryID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Delivery tracking not found"})
		return
	}

	c.JSON(http.StatusOK, deliveryTracking)
}
