package main

import (
	"log"
	"net/http"
	"time"

	"github.com/profclems/go-dotenv"
	"scada-lts.org/telegramsvc/handler"
)

func main() {
	// .env - It will search for the .env file in the current directory and load it.
	// You can explicitly set config file with dotenv.SetConfigFile("path/to/file.env")
	if err := dotenv.LoadConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Endpoints
	router := http.NewServeMux()
	router.HandleFunc("/webhook", handler.HandleMessage)

	// Static content
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/", fs)

	server := http.Server{
		Addr:         ":8088",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
