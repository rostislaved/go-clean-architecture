package config

import (
	httpAdapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter"
	kafka_queue "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber/kafka-queue"
	nats_adapter_subscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/nats-adapter-subscriber"
	pprofAdapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/pprof-adapter"
	books_gateway "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/gateways/books-gateway"
	kafka_adapter_publisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/kafka-adapter-publisher"
	nats_adapter_publisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/nats-adapter-publisher"
	books_repository_clickhouse "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/books-repository-clickhouse"
	books_repository_mongo "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/books-repository-mongo"
	books_repository_postgres "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/books-repository-postgres"
)

type Adapters struct {
	Primary   Primary
	Secondary Secondary
}

type Primary struct {
	HttpAdapter            httpAdapter.Config
	PprofAdapter           pprofAdapter.Config
	NatsAdapterSubscriber  nats_adapter_subscriber.Config
	KafkaAdapterSubscriber kafka_queue.Config
}

type Secondary struct {
	NatsAdapterPublisher  nats_adapter_publisher.Config
	KafkaAdapterPublisher kafka_adapter_publisher.Config
	Databases             Databases
	Gateways              Gateways
}

type Databases struct {
	Postgres   books_repository_postgres.Config
	Clickhouse books_repository_clickhouse.Config
	Mongo      books_repository_mongo.Config
}

type Gateways struct {
	Gateway1 books_gateway.Config
}
