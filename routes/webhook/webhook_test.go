package webhook

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/katesclau/telegramsvc/client"
)

func TestParseUpdateMessageWithText(t *testing.T) {
	chat := client.Chat{
		Id: 1,
	}

	var msg = client.Message{
		Text: "hello world",
		Chat: chat,
	}

	var update = client.Update{
		UpdateId: 1,
		Message:  msg,
	}

	requestBody, err := json.Marshal(update)
	if err != nil {
		t.Errorf("Failed to marshal update in json, got %s", err.Error())
	}
	req := httptest.NewRequest("POST", "http://localhost:8088/webhook", bytes.NewBuffer(requestBody))

	var updateToTest, errParse = parseTelegramRequest(req)
	if errParse != nil {
		t.Errorf("Expected a <nil> error, got %s", errParse.Error())
	}
	if *updateToTest != update {
		t.Errorf("Expected update %s, got %s", update, updateToTest)
	}

}
