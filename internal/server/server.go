package server

import (
	"context"
	"encoding/json"
	"log"

	"flux/pkg/interactions"

	"github.com/Haeven/codec/pkg/kafka"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Server struct {
	kafkaClient        *kafka.KafkaClient
	interactionService *interactions.InteractionService
}

func NewServer(kafkaClient *kafka.KafkaClient, interactionService *interactions.InteractionService) *Server {
	log.Println("Creating new server instance")
	return &Server{
		kafkaClient:        kafkaClient,
		interactionService: interactionService,
	}
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Starting server")

	// Start Kafka consumer for interactions
	go s.consumeInteractions(ctx)

	// ... (rest of the Run method remains unchanged)

	<-ctx.Done()
	return ctx.Err()
}

func (s *Server) consumeInteractions(ctx context.Context) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumerGroup("interaction_group"),
		kgo.ConsumeTopics("likes", "dislikes", "comments"),
	)
	if err != nil {
		log.Fatalf("Failed to create Kafka client: %v", err)
	}
	defer client.Close()

	for {
		fetches := client.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			for _, err := range errs {
				log.Printf("Kafka error: %v", err)
			}
			continue
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			switch record.Topic {
			case "likes":
				s.handleLike(record)
			case "dislikes":
				s.handleDislike(record)
			case "comments":
				s.handleComment(record)
			}
		}

		if ctx.Err() != nil {
			return
		}
	}
}

func (s *Server) handleLike(record *kgo.Record) {
	var data struct {
		VideoID string `json:"video_id"`
		UserID  string `json:"user_id"`
	}
	if err := json.Unmarshal(record.Value, &data); err != nil {
		log.Printf("Error unmarshaling like message: %v", err)
		return
	}

	if err := s.interactionService.HandleLike(data.VideoID, data.UserID); err != nil {
		log.Printf("Error handling like: %v", err)
	}
}

func (s *Server) handleDislike(record *kgo.Record) {
	var data struct {
		VideoID string `json:"video_id"`
		UserID  string `json:"user_id"`
	}
	if err := json.Unmarshal(record.Value, &data); err != nil {
		log.Printf("Error unmarshaling dislike message: %v", err)
		return
	}

	if err := s.interactionService.HandleDislike(data.VideoID, data.UserID); err != nil {
		log.Printf("Error handling dislike: %v", err)
	}
}

func (s *Server) handleComment(record *kgo.Record) {
	var data struct {
		VideoID string `json:"video_id"`
		UserID  string `json:"user_id"`
		Content string `json:"content"`
	}
	if err := json.Unmarshal(record.Value, &data); err != nil {
		log.Printf("Error unmarshaling comment message: %v", err)
		return
	}

	if err := s.interactionService.HandleComment(data.VideoID, data.UserID, data.Content); err != nil {
		log.Printf("Error handling comment: %v", err)
	}
}

// ... (rest of the file remains unchanged)
