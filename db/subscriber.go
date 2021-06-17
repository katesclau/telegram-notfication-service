package db

import (
	"log"

	"gorm.io/gorm"
)

type Subscriber struct {
	gorm.Model
	Channel string `gorm:"not null"`
	Enabled bool   `gorm:"default:true"`
	TopicID uint
	Topic   Topic
}

type SubscriberInput struct {
	Channel string
}

func (client *DBClient) GetTopicSubscribers(topicID uint) []Subscriber {
	subscribers := []Subscriber{}
	results := client.db.Select("ID", "Channel", "Enabled", "TopicID", "CreatedAt", "UpdatedAt").Where(&Subscriber{TopicID: topicID}).Find(&subscribers)
	if results.Error != nil {
		log.Printf("Failed to retrieve Topics: %s \n", results.Error)
	}
	return subscribers
}

func (client *DBClient) AddTopicSubscribers(topic Topic, input []SubscriberInput) {
	subscribers := []Subscriber{}
	for _, subscriber := range input {
		subscribers = append(subscribers, Subscriber{
			Channel: subscriber.Channel,
			TopicID: topic.ID,
		})
	}

	results := client.db.Create(subscribers)
	if results.Error != nil {
		log.Printf("Failed to create Subscribers: %s \n", results.Error)
	}

	if results.RowsAffected == 0 {
		log.Println("No subscribers created!")
	}
}
