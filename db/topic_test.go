package db

import (
	"log"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/katesclau/telegramsvc/utils"
	"github.com/profclems/go-dotenv"
	"github.com/stretchr/testify/assert"
)

func TestTopic(t *testing.T) {
	gofakeit.Seed(0)
	path, _ := os.Getwd()
	dotenv.SetConfigFile(utils.BuildString(path, "/../.env"))
	if err := dotenv.LoadConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var aSubscriber = SubscriberInput{
		Channel: gofakeit.LoremIpsumWord(),
	}
	var someSubscribers = []SubscriberInput{aSubscriber}

	var aTopicWithoutSubscribers = Topic{
		Name:        gofakeit.LoremIpsumWord(),
		Subscribers: []Subscriber{},
	}
	var aTopicWithSubscribers = Topic{
		Name:        gofakeit.LoremIpsumWord(),
		Subscribers: []Subscriber{{Channel: aSubscriber.Channel}},
	}
	// createdTopic := &Topic{}

	// Init DB
	dbclient := GetInstance(
		"MYSQL",                            // DB Type
		dotenv.GetString("MYSQL_URL"),      // Endpoint
		dotenv.GetString("MYSQL_DATABASE"), // Database
		dotenv.GetString("MYSQL_USER"),
		dotenv.GetString("MYSQL_PASSWORD"),
	)

	t.Run("Create a Topic", func(t *testing.T) {
		type args struct {
			name        string
			subscribers []SubscriberInput
		}
		tests := []struct {
			name   string
			client *DBClient
			args   args
			want   *Topic
		}{
			{
				"Create topic with no subscribers",
				dbclient,
				args{
					aTopicWithoutSubscribers.Name,
					[]SubscriberInput{},
				},
				&aTopicWithoutSubscribers,
			},
			{
				"Create topic with subscribers",
				dbclient,
				args{
					aTopicWithSubscribers.Name,
					someSubscribers,
				},
				&aTopicWithSubscribers,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.client.CreateTopic(tt.args.name, tt.args.subscribers)
				assert.Equal(t, got.Name, tt.want.Name, "Created %s Topic, and expected %s", got.Name, tt.want.Name)
				assert.Len(t, got.Subscribers, len(tt.want.Subscribers), "Created Topic with %d, and expected %d subscribers", len(got.Subscribers), len(tt.want.Subscribers))
			})
		}
	})

	t.Run("Get Topic", func(t *testing.T) {
		type args struct {
			name string
		}
		tests := []struct {
			name   string
			client *DBClient
			args   args
			want   *Topic
		}{
			{
				"Get topic with no subscribers",
				dbclient,
				args{
					aTopicWithoutSubscribers.Name,
				},
				&aTopicWithoutSubscribers,
			},
			{
				"Get topic unexisting topic",
				dbclient,
				args{
					gofakeit.LoremIpsumWord(),
				},
				nil,
			},
			{
				"Get topic with subscribers",
				dbclient,
				args{
					aTopicWithSubscribers.Name,
				},
				&aTopicWithSubscribers,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.client.GetTopic(tt.args.name)
				if tt.want == nil {
					assert.Nilf(t, got, "Got something, when expecting nil %v", got)
				}
				assert.Equal(t, got.Name, tt.want.Name, "Got %s Topic, and expected %s", got.Name, tt.want.Name)
				assert.Len(t, got.Subscribers, len(tt.want.Subscribers), "Got Topic with %d, and expected %d subscribers", len(got.Subscribers), len(tt.want.Subscribers))
			})
		}
	})

	t.Run("Get topics", func(t *testing.T) {
		tests := []struct {
			name   string
			client *DBClient
			want   []Topic
		}{
			{
				"Get All Topics",
				dbclient,
				[]Topic{aTopicWithoutSubscribers, aTopicWithoutSubscribers},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.client.GetTopics()
				assert.GreaterOrEqual(t, len(got), len(tt.want), "Received less Topics than expected!")
			})
		}
	})

	t.Run("Delete Topics", func(t *testing.T) {
		type args struct {
			name string
		}
		tests := []struct {
			name   string
			client *DBClient
			args   args
			want   *Topic
		}{
			{
				"Delete topic without subscribers",
				dbclient,
				args{
					aTopicWithoutSubscribers.Name,
				},
				&aTopicWithoutSubscribers,
			},
			{
				"Delete topic with subscribers",
				dbclient,
				args{
					aTopicWithSubscribers.Name,
				},
				&aTopicWithSubscribers,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.client.DeleteTopic(tt.args.name)
				assert.Equal(t, got.Name, tt.want.Name, "Deleted Topic do not match expected Name!")
			})
		}
	})
}
