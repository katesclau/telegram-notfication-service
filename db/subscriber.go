package db

import (
	"time"

	"gorm.io/gorm"
)

type Subscriber struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Channel   string         `gorm:"not null"`
	Enabled   bool           `gorm:"default:true"`
}
