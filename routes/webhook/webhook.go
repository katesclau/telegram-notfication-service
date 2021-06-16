/*
Package webhook implements a simple HTTP Request Handler for Telegram webhooks.
*/
package webhook

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/katesclau/telegramsvc/client"
	"github.com/katesclau/telegramsvc/db"
)

var DB *db.DBClient

func GetMethods(db *db.DBClient) map[string]func(w http.ResponseWriter, r *http.Request) {
	DB = db
	methods := make(map[string]func(w http.ResponseWriter, r *http.Request))
	methods["POST"] = HandleMessage
	return methods
}

// HandleMessage receives a Telegram webhook request,
// and responds through Telegram's API (see Package client)
func HandleMessage(w http.ResponseWriter, r *http.Request) {
	// Parse
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	// Sanitize input
	// var sanitizedSeed = sanitize(update.Message.Text)

	// Send the punchline back to Telegram
	client.TelegramClient.SendMessage(update.Message.Chat.Id, "Some data")
}

// parseTelegramRequest deserializes the JSON received in the request body
// from Telegram webhook request
func parseTelegramRequest(r *http.Request) (*client.Update, error) {
	var update client.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}
