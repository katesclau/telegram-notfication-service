package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/routes"
	"github.com/profclems/go-dotenv"
)

func main() {
	// .env - It will search for the .env file in the current directory and load it.
	// You can explicitly set config file with dotenv.SetConfigFile("path/to/file.env")
	if err := dotenv.LoadConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Init DB
	db := db.GetInstance(
		"MYSQL",                            // DB Type
		dotenv.GetString("MYSQL_URL"),      // Endpoint
		dotenv.GetString("MYSQL_DATABASE"), // Database
		dotenv.GetString("MYSQL_USER"),
		dotenv.GetString("MYSQL_PASSWORD"),
	)
	// Create a new Routes instance with the DB context - perhaps create a context struct later 😉
	routes := routes.NewRoutes(db)
	// We shouldn't run this everytime, only when models change
	db.AutoMigrate()

	servingErr := http.ListenAndServe(fmt.Sprintf(":%s", dotenv.GetString("PORT")), routes.GetRouter())
	if servingErr != nil {
		log.Fatal("Failed to init Server: %w", servingErr.Error())
	}
}
