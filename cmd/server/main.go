package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"WebAPIGo/config"
	"WebAPIGo/internal/handler"
	"WebAPIGo/internal/kafka"
	"WebAPIGo/internal/repository"
	"WebAPIGo/internal/server"
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
	defer func(producer *kafka.Producer) {
		err := producer.Close()
		if err != nil {

		}
	}(producer)

	router := gin.Default()

	repo := repository.NewPaymentRepository()
	svc := service.NewPaymentService(repo, producer)
	h := handler.NewPaymentHandler(svc)

	router.POST("/payment", h.HandlePayment)

	srv := server.NewServer(router, ":8080")

	go srv.Start()

	gracefulTimeout := 5 * time.Second // Set timeout for graceful shutdown.
	srv.GracefulShutdown(gracefulTimeout)

	log.Println("Application stopped.")
}
