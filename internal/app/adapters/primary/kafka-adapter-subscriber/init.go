package kafkaAdapterSubscriber

import (
	"log/slog"

	kafkaController "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber/kafka-controller"
	kafkaQueue "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber/kafka-queue"
	"github.com/rostislaved/go-clean-architecture/internal/app/application/service"
)

type KafkaAdapter struct {
	logger          *slog.Logger
	config          KafkaAdapterSubscriberConfig
	kafkaQueue      *kafkaQueue.KafkaQueue
	kafkaController *kafkaController.KafkaController
}

func New(l *slog.Logger, config KafkaAdapterSubscriberConfig, svc *service.ApiService) KafkaAdapter {
	kafkaQueue := kafkaQueue.New(l, config)

	kafkaController := kafkaController.New(l, svc)

	return KafkaAdapter{
		logger:          l,
		config:          config,
		kafkaQueue:      kafkaQueue,
		kafkaController: kafkaController,
	}
}
