package db

import (
	"log"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name        string `gorm:"not null:unique"`
	Subscribers []Subscriber
}
type TopicInput struct {
	Name        string
	Subscribers []SubscriberInput
}

func (topic *Topic) BeforeDelete(tx *gorm.DB) (err error) {
	var subscribers = []Subscriber{}
	results := tx.Where(&Subscriber{TopicID: topic.ID}).Find(&subscribers)
	if results.Error != nil {
		log.Printf("Failed to delete Topic: %s \n", results.Error)
	}

	tx.Delete(&subscribers)
	return
}

func (client *DBClient) CreateTopic(name string, subscribers []SubscriberInput) *Topic {
	topic := &Topic{Name: name}
	client.db.FirstOrCreate(topic, Topic{Name: name})
	client.AddTopicSubscribers(*topic, subscribers)
	topic = client.GetTopic(name)
	return topic
}

func (client *DBClient) GetTopic(name string) *Topic {
	var topic *Topic
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
	topic.Subscribers = client.GetTopicSubscribers(topic.ID)
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
