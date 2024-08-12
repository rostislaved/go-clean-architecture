package natsAdapterPublisher

import (
	"github.com/rostislaved/go-clean-architecture/internal/app/config"
)

type NatsAdapterPublisherConfig struct {
	Connection config.Connection

	Publisher1 Publisher
}
type Publisher struct {
	Channel string
}
