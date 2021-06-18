/*
Package webhook implements a simple HTTP Request Handler for Telegram webhooks.
*/
package webhook

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/katesclau/telegramsvc/routes/route"
	"github.com/katesclau/telegramsvc/telegram"
)

func GetRoute(ctx *route.Context) route.Route {
	r := route.NewRoute(ctx, "/webhook", false)
	methodHandlers := make(map[string]func(ctx *route.Context, w http.ResponseWriter, r *http.Request))
	methodHandlers["POST"] = postHandler
	r.SetMethodHandlers(methodHandlers)
	return *r
}

// Receives a Telegram webhook request,
// and responds through Telegram's API (see package telegram)
func postHandler(ctx *route.Context, w http.ResponseWriter, r *http.Request) {
	// Parse
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	ctx.WG.Add(1)
	go func() {
		defer ctx.WG.Done()
		// Sanitize input
		// var sanitizedSeed = sanitize(update.Message.Text)

		// Send the punchline back to Telegram
		telegram.TelegramClient.SendMessage(update.Message.Chat.Id, "Some data")
	}()
}

// parseTelegramRequest deserializes the JSON received in the request body
// from Telegram webhook request
func parseTelegramRequest(r *http.Request) (*telegram.Update, error) {
	var update telegram.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}
