package db

import (
	"fmt"
	"time"
)

type Topic struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `gorm:"not null:unique"`
	Subscribers []Subscriber
}

func (client *DBClient) CreateTopic(name string, subscribers []Subscriber) *Topic {
	topic := &Topic{Name: name, Subscribers: subscribers}
	client.db.FirstOrCreate(&topic, Topic{Name: name})
	return topic
}

func (client *DBClient) GetTopic(name string) *Topic {
	var topic *Topic
	client.db.First(&topic, Topic{Name: name})
	return topic
}

func (client *DBClient) GetTopics(name string) []Topic {
	var topics []Topic
	results := client.db.Find(&topics)
	if results.Error != nil {
		fmt.Printf("Failed to retrieve Topics: %s \n", results.Error)
	}
	return topics
}

func (client *DBClient) DeleteTopic(name string) *Topic {
	var topic *Topic
	client.db.Where("Name = ?", name).Delete(&topic)
	return topic
}
