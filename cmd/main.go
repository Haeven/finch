package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"flux/internal/server"
	"flux/pkg/account"
	"flux/pkg/interactions"

	"github.com/Haeven/codec/pkg/kafka"
)

func main() {
	// Create a new context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Initialize Kafka client
	kafkaClient, err := kafka.NewKafkaClient("localhost:9092")
	if err != nil {
		log.Fatalf("Failed to create Kafka client: %v", err)
	}
	defer kafkaClient.Close()

	// Initialize InteractionService
	interactionService := interactions.NewInteractionService()

	// Initialize AccountService
	accountService := account.NewAccountService()

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
