package models

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	RestaurantID uint           `gorm:"not null"`
	UserID       uint           `gorm:"not null"`
	Rating       int            `gorm:"not null"` // 1 to 5 stars
	Comment      string         `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
