package models

import (
	"gorm.io/gorm"
	"time"
)

type Menu struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	RestaurantID uint           `gorm:"not null"`
	Items        []MenuItem     `gorm:"foreignKey:MenuID"`
	CreatedAt    time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type MenuItem struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	MenuID      uint           `gorm:"not null"`
	Name        string         `gorm:"type:varchar(100);not null"`
	Description string         `gorm:"type:varchar(255);not null"`
	Price       float64        `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
