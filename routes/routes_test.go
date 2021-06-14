package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/katesclau/telegramsvc/db"
	"github.com/stretchr/testify/assert"
)

/**
Testing Topic Routes
*/
func TestTopics(t *testing.T) {
	var someSubscribers = []db.Subscriber{{
		Channel: "channel_id",
		Enabled: true,
	}}

	var aTopic = db.Topic{
		Name:        "SomeTopic",
		Subscribers: someSubscribers,
	}
	createdTopic := &db.Topic{}

	// Test Post Topic
	t.Run("Create a Topic", func(t *testing.T) {
		jsonTopic, err := json.Marshal(aTopic)
		if err != nil {
			t.FailNow()
		}
		request, _ := http.NewRequest("Post", "/topic", bytes.NewBuffer(jsonTopic))
		response := httptest.NewRecorder()
		router := GetRoutes()
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
		router := GetRoutes()
		router.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")
	})

	// Test Get a Topic
	t.Run("Get Created Topic", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/topic", nil)
		response := httptest.NewRecorder()
		router := GetRoutes()
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
		router := GetRoutes()
		router.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")
	})

}
