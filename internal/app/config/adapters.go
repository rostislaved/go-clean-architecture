package config

import (
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/nats-adapter-subscriber"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/gateways/gateway1"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/kafka-adapter-publisher"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/nats-adapter-publisher"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/booksRepository"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/documents-repository"
)

type Adapters struct {
	Primary   Primary
	Secondary Secondary
}

type Primary struct {
	HttpAdapter            httpAdapter.HttpAdapterConfig
	NatsAdapterSubscriber  natsAdapterSubscriber.NatsAdapterSubscriberConfig
	KafkaAdapterSubscriber kafkaAdapterSubscriber.KafkaAdapterSubscriberConfig
}

type Secondary struct {
	NatsAdapterPublisher  natsAdapterPublisher.NatsAdapterPublisherConfig
	KafkaAdapterPublisher kafkaAdapterPublisher.KafkaAdapterPublisherConfig
	Databases             Databases
	Gateways              Gateways
}
type Databases struct {
	Postgres   booksRepository.DatabaseRelational
	Clickhouse booksRepository.DatabaseRelational
	Mongo      employeesRepository.DatabaseMongo
}

type Gateways struct {
	Gateway1 gateway1.Config
}
