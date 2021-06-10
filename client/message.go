package client

import "fmt"

// Message is a Telegram object that can be found in an update.
type Message struct {
	MessageId int    `json:"messageId"`
	Text      string `json:"text"`
	Chat      Chat   `json:"chat"`
}

func (m Message) String() string {
	return fmt.Sprintf("{ Text: %s, Chat: %s }", m.Text, m.Chat)
}
