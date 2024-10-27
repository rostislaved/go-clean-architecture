package kafka_queue

import (
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaQueue struct {
	logger      *slog.Logger
	config      Config
	kafkaReader *kafka.Reader
}

type Config struct {
	Host    string
	GroupID string
	Topic   string
}

func New(
	l *slog.Logger,
	cfg Config,
) *KafkaQueue {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{cfg.Host},
		GroupID:  cfg.GroupID,
		Topic:    cfg.Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  1 * time.Second,
	})

	return &KafkaQueue{
		logger:      l,
		config:      cfg,
		kafkaReader: r,
	}
}
