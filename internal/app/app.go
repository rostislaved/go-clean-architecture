package app

import (
	"log/slog"

	http_adapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter"
	kafka_adapter_subscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber"
	nats_adapter_subscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/nats-adapter-subscriber"
	pprof_adapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/pprof-adapter"
	books_gateway "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/gateways/books-gateway"
	kafka_adapter_publisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/kafka-adapter-publisher"
	nats_adapter_publisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/nats-adapter-publisher"
	books_repository_postgres "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/books-repository-postgres"
	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
	"github.com/rostislaved/go-clean-architecture/internal/app/config"
)

type App struct {
	HttpAdapter            *http_adapter.HttpAdapter
	PprofAdapter           *pprof_adapter.PprofAdapter
	NatsAdapterSubscriber  *nats_adapter_subscriber.NatsAdapterSubscriber
	KafkaAdapterSubscriber *kafka_adapter_subscriber.KafkaAdapter
}

func New(l *slog.Logger, cfg config.Config) App {
	booksRepository := books_repository_postgres.New(l, cfg.Adapters.Secondary.Databases.Postgres)
	gateway := books_gateway.New(l, cfg.Adapters.Secondary.Gateways.Gateway1)
	natsAdapterPublisher := nats_adapter_publisher.New(l, cfg.Adapters.Secondary.NatsAdapterPublisher)
	kafkaAdapterPublisher := kafka_adapter_publisher.New(l, cfg.Adapters.Secondary.KafkaAdapterPublisher)

	usecases := usecases.New(
		l,
		cfg.Application.UseCases,
		booksRepository,
		gateway,
		natsAdapterPublisher,
		kafkaAdapterPublisher,
	)

	natsAdapterSubscriber := nats_adapter_subscriber.New(l, cfg.Adapters.Primary.NatsAdapterSubscriber, usecases)
	kafkaAdapter := kafka_adapter_subscriber.New(l, cfg.Adapters.Primary.KafkaAdapterSubscriber, usecases)
	httpAdapter := http_adapter.New(l, cfg.Adapters.Primary.HttpAdapter, usecases)
	pprofAdapter := pprof_adapter.New(l, cfg.Adapters.Primary.PprofAdapter)

	return App{
		HttpAdapter:            httpAdapter,
		PprofAdapter:           pprofAdapter,
		NatsAdapterSubscriber:  natsAdapterSubscriber,
		KafkaAdapterSubscriber: kafkaAdapter,
	}
}
