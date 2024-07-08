package models

import (
	"gorm.io/gorm"
	"time"
)

type Restaurant struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"type:varchar(100);not null"`
	Address     string         `gorm:"type:varchar(255);not null"`
	Phone       string         `gorm:"type:varchar(15);not null"`
	Email       string         `gorm:"type:varchar(100);not null"`
	Description string         `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
