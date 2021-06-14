package db

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"not null:unique"`
	Subscribers []Subscriber
}

func (client *DBClient) CreateTopic(name string, subscribers []Subscriber) *Topic {
	topic := &Topic{Name: name, Subscribers: subscribers}
	client.db.FirstOrCreate(&topic, Topic{Name: name})
	return topic
}
