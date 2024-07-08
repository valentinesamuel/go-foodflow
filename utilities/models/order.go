package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	UserID       uint           `gorm:"not null"`
	RestaurantID uint           `gorm:"not null"`
	Items        []OrderItem    `gorm:"foreignKey:OrderID"`
	TotalPrice   float64        `gorm:"not null"`
	Status       string         `gorm:"type:varchar(50);not null"`
	CreatedAt    time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type OrderItem struct {
	ID         uint           `gorm:"primaryKey;autoIncrement"`
	OrderID    uint           `gorm:"not null"`
	MenuItemID uint           `gorm:"not null"`
	Quantity   int            `gorm:"not null"`
	Price      float64        `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt  time.Time      `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
