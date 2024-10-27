package nats_adapter_publisher

import (
	"log/slog"
)

type NatsAdapterPublisher struct {
	logger    *slog.Logger
	config    Config
	publisher publisher
}

type publisher interface {
	Publish(channel string, data []byte) error
}

func New(logger *slog.Logger, config Config) *NatsAdapterPublisher {
	return &NatsAdapterPublisher{
		logger: logger,
		config: config,
	}
}
