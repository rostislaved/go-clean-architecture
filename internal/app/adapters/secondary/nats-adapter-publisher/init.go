package natsAdapterPublisher

import (
	"log/slog"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/config"
)

type NatsAdapterPublisher struct {
	logger    *slog.Logger
	config    config.NatsAdapterPublisher
	publisher publisher
}

type publisher interface {
	Publish(channel string, data []byte) error
}

func New(logger *slog.Logger, config config.NatsAdapterPublisher) *NatsAdapterPublisher {
	return &NatsAdapterPublisher{
		logger: logger,
		config: config,
	}
}
