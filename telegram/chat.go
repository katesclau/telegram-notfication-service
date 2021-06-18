package telegram

import "fmt"

// A Telegram Chat indicates the conversation to which the message belongs.
type Chat struct {
	Id int `json:"id"`
}

func (c Chat) String() string {
	return fmt.Sprintf("{ Id: %d }", c.Id)
}
