package db

import (
	"log"
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
	var subscribers []Subscriber
	topicResult := client.db.First(&topic, Topic{Name: name}).Joins("Subscriber")
	if topicResult.Error != nil {
		log.Println("Failed to retrieve topic", name, topicResult.Error)
		return nil
	}

	if topicResult.RowsAffected == 0 {
		log.Println("Topic not found ", name)
		return nil
	}

	// TODO: Use Joins https://gorm.io/docs/query.html#Joins-Preloading
	result := client.db.Where(&Subscriber{TopicID: topic.ID}).Find(&subscribers)
	if result.Error != nil {
		log.Println("Failed to include subscribers", result.Error)
		return topic
	}
	topic.Subscribers = subscribers
	return topic
}

func (client *DBClient) GetTopics() []Topic {
	var topics []Topic
	results := client.db.Find(&topics)
	if results.Error != nil {
		log.Printf("Failed to retrieve Topics: %s \n", results.Error)
	}
	log.Printf("Topics: %v", topics)
	return topics
}

func (client *DBClient) DeleteTopic(name string) *Topic {
	var topic *Topic
	client.db.Where("Name = ?", name).Delete(&topic)
	return topic
}
