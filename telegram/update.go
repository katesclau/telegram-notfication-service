package telegram

import "fmt"

// Update is a Telegram object that the handler receives every time an user interacts with the bot.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

func (u Update) String() string {
	return fmt.Sprintf("{ UpdateId: %d, Message: %s }", u.UpdateId, u.Message)
}
