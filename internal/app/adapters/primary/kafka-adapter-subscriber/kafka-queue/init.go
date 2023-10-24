package kafkaQueue

import (
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/config"
)

type KafkaQueue struct {
	logger      *slog.Logger
	config      config.KafkaAdapterSubscriber
	kafkaReader *kafka.Reader
}

func New(
	l *slog.Logger,
	cfg config.KafkaAdapterSubscriber,
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
