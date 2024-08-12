package natsAdapterPublisher

import (
	"log/slog"
)

type NatsAdapterPublisher struct {
	logger    *slog.Logger
	config    NatsAdapterPublisherConfig
	publisher publisher
}

type publisher interface {
	Publish(channel string, data []byte) error
}

func New(logger *slog.Logger, config NatsAdapterPublisherConfig) *NatsAdapterPublisher {
	return &NatsAdapterPublisher{
		logger: logger,
		config: config,
	}
}
