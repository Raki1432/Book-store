package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rakesh-honawad/Book-store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Register the routes
	routes.RegisterBookStoreRoutes(r)

	// Setup the server
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:9010",
	}

	// Run the server in a goroutine
	go func() {
		log.Println("Starting server on :9010")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// Setup graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a shutdown signal
	sigReceived := <-c
	log.Printf("Received signal: %v, shutting down...", sigReceived)

	// Set a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*1000000000) // 5 seconds
	defer cancel()

	// Gracefully shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
