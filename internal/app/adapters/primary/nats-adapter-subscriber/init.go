package nats_adapter_subscriber

import (
	"log/slog"

	natsController "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/nats-adapter-subscriber/nats-handlers"
	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
)

type NatsAdapterSubscriber struct {
	logger         *slog.Logger
	config         Config
	subscriber     subscriber
	svc            *usecases.UseCases
	natsController *natsController.NatsHandlers
}

type subscriber interface {
	// Subscribe(cfg SubscriptionConfig) (*Subscription, error)
}

func New(logger *slog.Logger, config Config, svc *usecases.UseCases) *NatsAdapterSubscriber {
	natsController := natsController.New(logger, svc)

	return &NatsAdapterSubscriber{
		logger: logger,
		config: config,
		// subscriber:     a,
		svc:            svc,
		natsController: natsController,
	}
}
