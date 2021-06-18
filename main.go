package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/routes"
	"github.com/katesclau/telegramsvc/routes/route"
	"github.com/katesclau/telegramsvc/telegram"
	"github.com/profclems/go-dotenv"
)

func main() {
	// .env - It will search for the .env file in the current directory and load it.
	// You can explicitly set config file with dotenv.SetConfigFile("path/to/file.env")
	if err := dotenv.LoadConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Sync group for gracefully shutdown
	wg := &sync.WaitGroup{}

	// Init DB
	db := db.GetInstance(
		"MYSQL",                            // DB Type
		dotenv.GetString("MYSQL_URL"),      // Endpoint
		dotenv.GetString("MYSQL_DATABASE"), // Database
		dotenv.GetString("MYSQL_USER"),
		dotenv.GetString("MYSQL_PASSWORD"),
	)
	// Create a new Routes instance with the DB context - perhaps create a context struct later 😉

	ctx := &route.Context{
		DB:             db,
		WG:             wg,
		TelegramClient: telegram.TelegramClient,
	}

	routes := routes.NewRoutes(ctx)
	// We shouldn't run this everytime, only when models change
	db.AutoMigrate()

	router := routes.GetRouter()
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", dotenv.GetString("PORT")),
		Handler: router,
	}
	servingErr := httpServer.ListenAndServe()
	if servingErr != nil {
		log.Fatal("Failed to init Server: %w", servingErr.Error())
	}

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-termChan // Blocks here until interrupted
		log.Print("SIGTERM received. Shutdown process initiated\n")
		httpServer.Shutdown(context.Background())
	}()

	// This is where, once we're closing the program, we wait for all
	// jobs (they all have been added to this WaitGroup) to `wg.Done()`.
	log.Println("waiting for running jobs to finish")
	wg.Wait()
	log.Println("jobs finished. exiting")
}

//
