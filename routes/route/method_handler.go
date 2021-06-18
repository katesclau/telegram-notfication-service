package route

import (
	"net/http"
	"sync"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/telegram"
)

type Context struct {
	DB             *db.DBClient
	TelegramClient *telegram.Client
	WG             *sync.WaitGroup
}

type MethodHandler struct {
	Method  string
	handler func(c *Context, w http.ResponseWriter, r *http.Request)
	context *Context
}

func (methodHandler *MethodHandler) Execute(w http.ResponseWriter, r *http.Request) {
	methodHandler.handler(methodHandler.context, w, r)
}

func NewMethodHandler(c *Context, m string, h func(c *Context, w http.ResponseWriter, r *http.Request)) *MethodHandler {
	methodHandler := &MethodHandler{
		Method:  m,
		handler: h,
		context: c,
	}
	return methodHandler
}
