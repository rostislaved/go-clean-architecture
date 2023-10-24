package main

import (
	"log/slog"
	"os"

	_ "go.uber.org/automaxprocs"

	httpAdapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/gateways/gateway1"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/firstRepository"

	osSignalAdapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/os-signal-adapter"

	natsAdapterSubscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/nats-adapter-subscriber"
	natsAdapterPublisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/nats-adapter-publisher"

	kafkaAdapterSubscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber"
	kafkaAdapterPublisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/kafka-adapter-publisher"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/service"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/config"
)

func main() {
	cfg := config.New()

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	l := slog.New(h)
	//// SECONDARY ADAPTERS

	// Repository adapter
	firstRepository := firstRepository.New(l, cfg.Adapters.Secondary.Databases.Postgres)
	_ = firstRepository // Используем в сервисе

	// Provider adapter
	provider := gateway1.New(l, cfg.Adapters.Secondary.Gateways.Gateway1)
	_ = provider // Используем в сервисе

	// Nats Adapter Publisher
	natsAdapterPublisher := natsAdapterPublisher.New(l, cfg.Adapters.Secondary.NatsAdapterPublisher)
	_ = natsAdapterPublisher // Используем в сервисе

	// Kafka Adapter Publisher
	kafkaAdapterPublisher := kafkaAdapterPublisher.New(l, cfg.Adapters.Secondary.KafkaAdapterPublisher)
	_ = kafkaAdapterPublisher // Используем в сервисе

	// legacy которое надо переписать
	// prometheus
	// health checks

	//// APPLICATION
	svc := service.New(l, cfg.Services.UpdateService, firstRepository, provider)

	//// PRIMARY ADAPTERS

	// Nats Adapter Subscriber
	natsAdapterSubscriber := natsAdapterSubscriber.New(l, cfg.Adapters.Primary.NatsAdapterSubscriber, svc)

	go natsAdapterSubscriber.Start()

	// Kafka Adapter Subscriber
	kafkaAdapter := kafkaAdapterSubscriber.New(l, cfg.Adapters.Primary.KafkaAdapterSubscriber, svc)

	go kafkaAdapter.Start()

	// Http Adapter
	httpAdapter := httpAdapter.New(l, cfg.Adapters.Primary.HttpAdapter, svc)

	go httpAdapter.Start()

	// OS signal adapter
	osSignalAdapter := osSignalAdapter.New()

	go osSignalAdapter.Start()

	// Graceful shutdown
	select {
	case err := <-httpAdapter.Notify():
		l.Error(err.Error(), "main")
	case err := <-osSignalAdapter.Notify():
		l.Error(err.Error(), "main")
	}
}
