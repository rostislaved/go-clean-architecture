package natsAdapterSubscriber

import (
	"log/slog"

	natsController "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/nats-adapter-subscriber/nats-controller"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/service"
)

type NatsAdapterSubscriber struct {
	logger         *slog.Logger
	config         NatsAdapterSubscriberConfig
	subscriber     subscriber
	svc            *service.ApiService
	natsController *natsController.NatsController
}

type subscriber interface {
	// Subscribe(cfg SubscriptionConfig) (*Subscription, error)
}

func New(logger *slog.Logger, config NatsAdapterSubscriberConfig, svc *service.ApiService) *NatsAdapterSubscriber {
	natsController := natsController.New(logger, svc)

	return &NatsAdapterSubscriber{
		logger: logger,
		config: config,
		// subscriber:     a,
		svc:            svc,
		natsController: natsController,
	}
}
