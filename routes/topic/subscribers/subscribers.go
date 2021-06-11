package subscribers

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Switch Methods
	switch r.Method {
	case "GET":
		log.Println("Got HERE!!!")
		w.Write([]byte(http.StatusText(http.StatusOK)))
	default:
		log.Println("Method not supported!", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}
