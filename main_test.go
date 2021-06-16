package main

import (
	"log"
	"os"
	"testing"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/utils"
	"github.com/profclems/go-dotenv"
)

func Test_main(t *testing.T) {
	log.Println("Prepare DB...")

	path, _ := os.Getwd()
	dotenv.SetConfigFile(utils.BuildString(path, "/.env"))
	if err := dotenv.LoadConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Init DB
	dbclient := db.GetInstance(
		"MYSQL",                            // DB Type
		dotenv.GetString("MYSQL_URL"),      // Endpoint
		dotenv.GetString("MYSQL_DATABASE"), // Database
		dotenv.GetString("MYSQL_USER"),
		dotenv.GetString("MYSQL_PASSWORD"),
	)

	// Migrate Models
	dbclient.AutoMigrate()
}
