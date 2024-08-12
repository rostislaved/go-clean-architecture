package kafkaAdapterPublisher

import (
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type KafkaAdapterPublisher struct {
	logger *slog.Logger
	config KafkaAdapterPublisherConfig
	writer *kafka.Writer
}

func New(logger *slog.Logger, config KafkaAdapterPublisherConfig) *KafkaAdapterPublisher {
	w := kafka.Writer{
		// TODO Брать поля из конфига
	}

	return &KafkaAdapterPublisher{
		logger: logger,
		config: config,
		writer: &w,
	}
}
