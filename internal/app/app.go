package app

import (
	"log/slog"

	grpc_adapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/grpc-adapter"
	http_adapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter"
	kafka_adapter_subscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber"
	nats_adapter_subscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/nats-adapter-subscriber"
	pprof_adapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/pprof-adapter"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/gateways/entity5-gateway"
	kafka_adapter_publisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/kafka-adapter-publisher"
	nats_adapter_publisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/nats-adapter-publisher"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/entity1-repository"
	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
	"github.com/rostislaved/go-clean-architecture/internal/app/config"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/postgres"
)

type App struct {
	HttpAdapter            *http_adapter.HttpAdapter
	GrpcAdapter            *grpc_adapter.GrpcAdapter
	PprofAdapter           *pprof_adapter.PprofAdapter
	NatsAdapterSubscriber  *nats_adapter_subscriber.NatsAdapterSubscriber
	KafkaAdapterSubscriber *kafka_adapter_subscriber.KafkaAdapter
}

func New(l *slog.Logger, cfg config.Config) App {
	db, err := postgres.Pgx(l, cfg.Infrastructure.Databases.Postgres)
	if err != nil {
		panic(err)
	}

	entity1Repository := entity1_repository.New(l, cfg.Adapters.Secondary.Entity1Config, db)
	entity5Gateway := entity5_gateway.New(l, cfg.Adapters.Secondary.Gateways.Entity5Gateway)
	natsAdapterPublisher := nats_adapter_publisher.New(l, cfg.Adapters.Secondary.NatsAdapterPublisher)
	kafkaAdapterPublisher := kafka_adapter_publisher.New(l, cfg.Adapters.Secondary.KafkaAdapterPublisher)

	usecases := usecases.New(
		l,
		cfg.Application.UseCases,
		entity1Repository,
		entity5Gateway,
		natsAdapterPublisher,
		kafkaAdapterPublisher,
	)

	httpAdapter := http_adapter.New(l, cfg.Adapters.Primary.HttpAdapter, usecases)
	grpcAdapter := grpc_adapter.New()
	pprofAdapter := pprof_adapter.New(l, cfg.Adapters.Primary.PprofAdapter)
	natsAdapterSubscriber := nats_adapter_subscriber.New(l, cfg.Adapters.Primary.NatsAdapterSubscriber, usecases)
	kafkaAdapter := kafka_adapter_subscriber.New(l, cfg.Adapters.Primary.KafkaAdapterSubscriber, usecases)

	return App{
		HttpAdapter:            httpAdapter,
		GrpcAdapter:            grpcAdapter,
		PprofAdapter:           pprofAdapter,
		NatsAdapterSubscriber:  natsAdapterSubscriber,
		KafkaAdapterSubscriber: kafkaAdapter,
	}
}
