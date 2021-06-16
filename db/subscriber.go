package db

import (
	"log"
	"time"
)

type Subscriber struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Channel   string `gorm:"not null"`
	Enabled   bool   `gorm:"default:true"`
	Topic     Topic  `gorm:"constraint:OnDelete:CASCADE;"`
	TopicID   uint
}

func (client *DBClient) GetTopicSubscribers(topicName string) []Subscriber {
	topic := client.GetTopic(topicName)
	subscribers := []Subscriber{}
	results := client.db.Where(&Subscriber{TopicID: topic.ID}).Find(&subscribers)
	if results.Error != nil {
		log.Printf("Failed to retrieve Topics: %s \n", results.Error)
	}
	return subscribers
}
