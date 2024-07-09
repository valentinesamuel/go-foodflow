package service

import (
	"DeliveryService/db"
	"go-foodflow/utilities/models"
)

func CreateDeliveryTrackingService(deliveryTracking *models.DeliveryTracking) (models.DeliveryTracking, error) {
	err := db.DeliveryDB.Create(&deliveryTracking).Error
	return *deliveryTracking, err
}

func GetDeliveryTrackingStatusService(deliveryID uint) (models.DeliveryTracking, error) {
	var deliveryTracking models.DeliveryTracking
	err := db.DeliveryDB.Table("delivery_trackings").Where("delivery_id = ?", deliveryID).First(&deliveryTracking).Error
	return deliveryTracking, err
}
