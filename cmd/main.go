package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"finch/internal/db" // Add this import
	"finch/pkg/account"
	"finch/pkg/interactions"

	socketio "github.com/googollee/go-socket.io"

	"finch/internal/server"

	"github.com/Haeven/codec/pkg/kafka"
)

func main() {
	// Create a new context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Initialize database
	database := db.Initialize()
	defer database.Close()

	// Initialize Kafka client
	kafkaClient, err := kafka.NewKafkaClient([]string{"localhost:9092"}, "interaction_group", "likes", "dislikes")
	if err != nil {
		log.Fatalf("Failed to create Kafka client: %v", err)
	}
	defer kafkaClient.Close()

	// Initialize InteractionService
	socketServer := socketio.NewServer(nil)
	if socketServer == nil {
		log.Fatalf("Failed to create socket.io server")
	}
	interactionService := interactions.NewInteractionService(socketServer)

	// Initialize AccountService
	accountService := account.NewAccountService(database)
	// Create and run the server
	srv := server.NewServer(kafkaClient, interactionService, accountService)

	// Run the server in a goroutine
	go func() {
		if err := srv.Run(ctx); err != nil {
			log.Printf("Server error: %v", err)
			cancel()
		}
	}()

	// Wait for shutdown signal
	<-sigChan
	log.Println("Shutting down gracefully...")
	cancel()
}
