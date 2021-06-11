package topic

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Switch Methods
	switch r.Method {
	case "GET":
		w.Write([]byte(http.StatusText(200)))
	case "PUT":
		w.Write([]byte(http.StatusText(200)))
	case "POST":
		w.Write([]byte(http.StatusText(200)))
	case "DELETE":
		w.Write([]byte(http.StatusText(200)))
	}
}
