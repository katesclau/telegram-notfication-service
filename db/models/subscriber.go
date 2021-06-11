package models

import "gorm.io/gorm"

type Subscriber struct {
	gorm.Model
	Channel string `gorm:"not null"`
	Enabled bool   `gorm:"default:true"`
	Topics  *Topic
}
