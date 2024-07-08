package models

import (
	"gorm.io/gorm"
	"time"
)

type DeliveryTracking struct {
	ID                      uint           `gorm:"primaryKey;autoIncrement"`
	OrderID                 uint           `gorm:"not null"`
	CurrentStatus           string         `gorm:"type:varchar(50);not null"`
	EstimatedTimeOfDelivery time.Time      `gorm:"type:timestamp;not null"`
	CurrentLatitude         float64        `gorm:"not null"`
	CurrentLongitude        float64        `gorm:"not null"`
	CreatedAt               time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt               time.Time      `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt               gorm.DeletedAt `gorm:"index"`
}
