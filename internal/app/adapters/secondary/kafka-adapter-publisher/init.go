package kafkaAdapterPublisher

import (
	"github.com/segmentio/kafka-go"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/config"
)

type KafkaAdapterPublisher struct {
	logger *slog.Logger
	config config.KafkaAdapterPublisher
	writer *kafka.Writer
}

func New(logger *slog.Logger, config config.KafkaAdapterPublisher) *KafkaAdapterPublisher {
	w := kafka.Writer{
		// TODO Брать поля из конфига
	}

	return &KafkaAdapterPublisher{
		logger: logger,
		config: config,
		writer: &w,
	}
}
