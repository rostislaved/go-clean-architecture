package config

import "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/booksRepository"

type Adapters struct {
	Primary   Primary
	Secondary Secondary
}

type Primary struct {
	HttpAdapter            HttpAdapter
	NatsAdapterSubscriber  NatsAdapterSubscriber
	KafkaAdapterSubscriber KafkaAdapterSubscriber
}

type Secondary struct {
	NatsAdapterPublisher  NatsAdapterPublisher
	KafkaAdapterPublisher KafkaAdapterPublisher
	Databases             Databases
	Gateways              Gateways
}
type Databases struct {
	Postgres   booksRepository.DatabaseRelational
	Clickhouse booksRepository.DatabaseRelational
	Mongo      DatabaseMongo
}

type Gateways struct {
	Gateway1 Gateway1
}
