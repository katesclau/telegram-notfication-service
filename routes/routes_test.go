package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/routes/route"
	"github.com/katesclau/telegramsvc/telegram"
	"github.com/katesclau/telegramsvc/utils"
	"github.com/profclems/go-dotenv"
	"github.com/stretchr/testify/assert"
)

/**
Testing Topic Routes
*/
func TestTopics(t *testing.T) {
	path, _ := os.Getwd()
	dotenv.SetConfigFile(utils.BuildString(path, "/../.env"))
	if err := dotenv.LoadConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var someSubscribers = []db.Subscriber{{
		Channel: "channel_id",
		Enabled: true,
	}}

	var aTopic = db.Topic{
		Name:        "SomeTopic",
		Subscribers: someSubscribers,
	}
	createdTopic := &db.Topic{}

	// Sync group for gracefully shutdown
	wg := &sync.WaitGroup{}

	// Init DB
	db := db.GetInstance(
		"MYSQL",                            // DB Type
		dotenv.GetString("MYSQL_URL"),      // Endpoint
		dotenv.GetString("MYSQL_DATABASE"), // Database
		dotenv.GetString("MYSQL_USER"),
		dotenv.GetString("MYSQL_PASSWORD"),
	)

	ctx := &route.Context{
		DB:             db,
		WG:             wg,
		TelegramClient: telegram.TelegramClient,
	}

	routes := NewRoutes(ctx)
	router := routes.GetRouter()

	// Test Post Topic
	t.Run("Create a Topic", func(t *testing.T) {
		jsonTopic, err := json.Marshal(aTopic)
		if err != nil {
			t.FailNow()
		}
		request, _ := http.NewRequest("Post", "/topic", bytes.NewBuffer(jsonTopic))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusAccepted, response.Code, "OK response is expected")

		data, err := ioutil.ReadAll(response.Body)
		if err := json.Unmarshal(data, createdTopic); err != nil {
			t.FailNow()
		}
		assert.NotNil(t, createdTopic.ID, "Topic created correcly")
		assert.Equal(t, aTopic.Name, createdTopic.Name, "Returned the Topic with proper name")
	})

	// Test Get Topics
	t.Run("Get Created Topic", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/topic", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")
	})

	// Test Get a Topic
	t.Run("Get Created Topic", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/topic", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")
	})

	// Test Subscribers
	t.Run("Get Topic Subscribers", func(t *testing.T) {
		// Add a subscriber through DB

		// Get through endpoint

		// Assert subscriber info
		request, _ := http.NewRequest("GET", "/topic", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")
	})

	// Test Events
	t.Run("Post an Event to a Topic", func(t *testing.T) {
		// Add a subscriber through DB

		// Get through endpoint

		// Assert subscriber info
		request, _ := http.NewRequest("POST", utils.BuildString("/topic/", createdTopic.Name, "/event"), nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")
	})

	// Delete a Topic
	t.Run("Delete the created Topic", func(t *testing.T) {
		request, _ := http.NewRequest("DELETE", utils.BuildString("/topic/", createdTopic.Name), nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusNoContent, response.Code, "OK response is expected")
	})
}
