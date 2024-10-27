package app

import (
	"log/slog"

	httpAdapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter"
	kafkaAdapterSubscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber"
	natsAdapterSubscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/nats-adapter-subscriber"
	pprofAdapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/pprof-adapter"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/gateways/gateway1"
	kafkaAdapterPublisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/kafka-adapter-publisher"
	natsAdapterPublisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/nats-adapter-publisher"
	books_repository_postgres "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/books-repository-postgres"
	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
	"github.com/rostislaved/go-clean-architecture/internal/app/config"
)

type App struct {
	HttpAdapter            *httpAdapter.HttpAdapter
	PprofAdapter           *pprofAdapter.PprofAdapter
	NatsAdapterSubscriber  *natsAdapterSubscriber.NatsAdapterSubscriber
	KafkaAdapterSubscriber *kafkaAdapterSubscriber.KafkaAdapter
}

func New(l *slog.Logger, cfg config.Config) App {
	booksRepository := books_repository_postgres.New(l, cfg.Adapters.Secondary.Databases.Postgres)
	gateway := gateway1.New(l, cfg.Adapters.Secondary.Gateways.Gateway1)
	natsAdapterPublisher := natsAdapterPublisher.New(l, cfg.Adapters.Secondary.NatsAdapterPublisher)
	kafkaAdapterPublisher := kafkaAdapterPublisher.New(l, cfg.Adapters.Secondary.KafkaAdapterPublisher)

	usecases := usecases.New(
		l,
		cfg.Application.UseCases,
		booksRepository,
		gateway,
		natsAdapterPublisher,
		kafkaAdapterPublisher,
	)

	natsAdapterSubscriber := natsAdapterSubscriber.New(l, cfg.Adapters.Primary.NatsAdapterSubscriber, usecases)
	kafkaAdapter := kafkaAdapterSubscriber.New(l, cfg.Adapters.Primary.KafkaAdapterSubscriber, usecases)
	httpAdapter := httpAdapter.New(l, cfg.Adapters.Primary.HttpAdapter, usecases)
	pprofAdapter := pprofAdapter.New(l, cfg.Adapters.Primary.PprofAdapter)

	return App{
		HttpAdapter:            httpAdapter,
		PprofAdapter:           pprofAdapter,
		NatsAdapterSubscriber:  natsAdapterSubscriber,
		KafkaAdapterSubscriber: kafkaAdapter,
	}
}
