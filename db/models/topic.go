package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Subscribers *Subscriber
	Name        string `gorm:"not null"`
}
