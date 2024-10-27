package kafka_adapter_subscriber

import (
	"log/slog"

	kafka_handlers "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber/kafka-handlers"
	kafka_queue "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber/kafka-queue"
	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
)

type KafkaAdapter struct {
	logger          *slog.Logger
	config          kafka_queue.Config
	kafkaQueue      *kafka_queue.KafkaQueue
	kafkaController *kafka_handlers.KafkaHandlers
}

func New(l *slog.Logger, config kafka_queue.Config, svc *usecases.UseCases) *KafkaAdapter {
	kafkaQueue := kafka_queue.New(l, config)

	kafkaController := kafka_handlers.New(l, svc)

	return &KafkaAdapter{
		logger:          l,
		config:          config,
		kafkaQueue:      kafkaQueue,
		kafkaController: kafkaController,
	}
}
