package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/profclems/go-dotenv"
	"scada-lts.org/telegramsvc/routes"
)

func main() {
	// .env - It will search for the .env file in the current directory and load it.
	// You can explicitly set config file with dotenv.SetConfigFile("path/to/file.env")
	if err := dotenv.LoadConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	servingErr := http.ListenAndServe(fmt.Sprintf(":%s", dotenv.GetString("PORT")), routes.GetRoutes())
	if servingErr != nil {
		log.Fatal("Failed to init Server: %w", servingErr.Error())
	}
}
