package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	syncProducer sarama.SyncProducer
	topic        string
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{
		syncProducer: producer,
		topic:        topic,
	}, nil
}

func (p *Producer) SendMessage(key string, value string) error {
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	partition, offset, err := p.syncProducer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message sent to partition %d with offset %d\n", partition, offset)
	return nil
}

func (p *Producer) Close() error {
	return p.syncProducer.Close()
}
