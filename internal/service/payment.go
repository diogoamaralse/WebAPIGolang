package service

import (
	"context"
	"encoding/json"

	"WebAPIGo/internal/kafka"
	"WebAPIGo/internal/model"
	"WebAPIGo/internal/repository"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, req model.PaymentRequest) (model.PaymentResponse, error)
}

type paymentService struct {
	repo     repository.PaymentRepository
	producer *kafka.Producer
}

func NewPaymentService(repo repository.PaymentRepository, producer *kafka.Producer) PaymentService {
	return &paymentService{repo: repo, producer: producer}
}

func (s *paymentService) ProcessPayment(ctx context.Context, req model.PaymentRequest) (model.PaymentResponse, error) {
	id, err := s.repo.SavePayment(ctx, req)
	if err != nil {
		return model.PaymentResponse{}, err
	}

	// Serialize the payment request to JSON for Kafka message.
	messageData, _ := json.Marshal(req)

	// Send the message to Kafka.
	if err := s.producer.SendMessage(id, string(messageData)); err != nil {
		return model.PaymentResponse{}, err
	}

	return model.PaymentResponse{ID: id}, nil
}
