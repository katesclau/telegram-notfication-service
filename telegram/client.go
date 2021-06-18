package telegram

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/profclems/go-dotenv"
)

// TODO: Make this singleton
var TelegramClient *Client = NewClient(dotenv.GetString("TELEGRAM_TOKEN"))
var BaseURL string = "https://api.telegram.org"

type Client struct {
	token string
}

func (cl *Client) SendMessage(chatId int, text string) string {
	log.Println("Client::SendMessage: ", chatId, text)
	var res Message
	TelegramClient.Call("sendMessage", url.Values{
		"chat_id":    []string{strconv.Itoa(chatId)},
		"text":       []string{text},
		"parse_mode": []string{"MarkdownV2"},
	}, &res)
	return fmt.Sprintf("MessageId: %d", res.MessageId)
}

func (cl *Client) Call(method string, params url.Values, v interface{}) {
	apiUrl, _ := buildApiUrl(BaseURL, cl.Bot(), method)
	apiUrl.RawQuery = params.Encode()

	resp, errors := http.Get(apiUrl.String())
	if errors != nil {
		log.Printf("Failed to Call %s - %s", method, errors.Error())
	}
	defer resp.Body.Close()

	log.Println("Client::Call: ", resp.Body, resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(v)
}

func (cl *Client) GetMe() Account {
	resp := AccountResponse{}
	cl.Call("getMe", url.Values{}, &resp)
	return resp.Account
}

func (cl *Client) GetFile(fileId string) File {
	resp := FileResponse{}
	cl.Call("getFile", url.Values{
		"file_id": []string{fileId},
	}, &resp)

	return resp.File
}

func (cl *Client) ReadFile(fileId string) File {
	file := cl.GetFile(fileId)
	apiUrl, _ := buildApiUrl(BaseURL, "file", cl.Bot(), file.Path)
	resp, errors := http.Get(apiUrl.String())
	if errors != nil {
		log.Printf("Failed to ReadFile %s - %s", fileId, errors)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	file.Content = "data:image/png;base64," + base64.StdEncoding.EncodeToString(body)
	return file
}

func (cl *Client) Bot() string {
	return "bot" + cl.token
}

func NewClient(token string) *Client {
	return &Client{token}
}

func buildApiUrl(parts ...string) (*url.URL, error) {
	rawurl := strings.Join(parts, "/")
	return url.Parse(rawurl)
}
