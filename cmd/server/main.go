package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"WebAPIGo/config"
	"WebAPIGo/internal/handler"
	"WebAPIGo/internal/kafka"
	"WebAPIGo/internal/repository"
	"WebAPIGo/internal/service"
)

// Structuring Large Go Projects
// Follow Standard Go Project Layout (pkg/, cmd/, internal/).
// Separate concerns using repository, service, handler patterns.
//├── cmd/        # Main entry points
//├── internal/   # Private application code
//├── pkg/        # Reusable packages
//├── api/        # API definitions
//├── config/     # Configuration files
//├── docs/       # Documentation

func main() {
	// Load configuration.
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Kafka producer.
	producer, err := kafka.NewProducer(cfg.Kafka.Brokers, cfg.Kafka.Topic)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}
	defer producer.Close()

	// Initialize repository and service.
	repo := repository.NewPaymentRepository()
	svc := service.NewPaymentService(repo, producer)

	// Initialize handler.
	h := handler.NewPaymentHandler(svc)

	r := gin.Default()
	r.POST("/payment", h.HandlePayment)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
