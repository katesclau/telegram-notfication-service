package db

import (
	"time"
)

type Subscriber struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Channel   string `gorm:"not null"`
	Enabled   bool   `gorm:"default:true"`
	Topic     Topic
	TopicID   uint
}
