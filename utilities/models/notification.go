package models

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	UserID    uint           `gorm:"not null"`
	Message   string         `gorm:"type:varchar(255);not null"`
	Status    string         `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
