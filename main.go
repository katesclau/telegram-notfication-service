// Telegram Service is a notification service that allows bot owners to issue notifications to subscribers in Telegram Channels
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
	"time"

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

	waitForShutdown(httpServer, wg)
}

func waitForShutdown(srv *http.Server, wg *sync.WaitGroup) {
	// Handle sigterm and await termChan signal
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	go func() {
		<-interruptChan
		log.Print("SIGTERM received. Shutdown process initiated\n")
		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	// This is where, once we're closing the program, we wait for all
	// jobs (they all have been added to this WaitGroup) to `wg.Done()`.
	log.Println("waiting for running jobs to finish")
	wg.Wait()
	log.Println("jobs finished. exiting")
	os.Exit(0)
}
