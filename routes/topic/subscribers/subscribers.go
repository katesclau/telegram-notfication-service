package subscribers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/katesclau/telegramsvc/db"
)

func getTopicSubscribers(topicName string) ([]db.Subscriber, error) {
	return nil, nil
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	pathArray := strings.Split(path, "/")

	subscribers, err := getTopicSubscribers(pathArray[1])
	if err != nil {
		msg := fmt.Sprintf("Failed to retrieve subscribers. %s", err)
		fmt.Print(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}

	jsonBody, marshalingErr := json.Marshal(subscribers)
	if marshalingErr != nil {
		msg := fmt.Sprintf("Failed to retrieve subscribers. %s", marshalingErr)
		fmt.Print(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBody)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Switch Methods
	switch r.Method {
	case "GET":
		getHandler(w, r)
	default:
		log.Println("Method not supported!", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}
