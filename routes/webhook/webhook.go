package webhook

import (
	"encoding/json"
	"log"
	"net/http"

	"scada-lts.org/telegramsvc/client"
)

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

func parseTelegramRequest(r *http.Request) (*client.Update, error) {
	var update client.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}
