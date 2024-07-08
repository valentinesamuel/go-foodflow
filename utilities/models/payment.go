package models

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	OrderID   uint           `gorm:"not null"`
	UserID    uint           `gorm:"not null"`
	Amount    float64        `gorm:"not null"`
	Status    string         `gorm:"type:varchar(50);not null"`
	Method    string         `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
