package config

import (
	"time"
)

type Config struct {
	Application Application
	Services    Services
	Adapters    Adapters
}
type Application struct {
	Name    string
	Version string
}

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
	Postgres   DatabaseRelational
	Clickhouse DatabaseRelational
	Mongo      DatabaseMongo
}

type DatabaseRelational struct {
	Type       string
	Host       string `config:"envVar"`
	Port       string `config:"envVar"`
	User       string `config:"envVar"`
	Password   string `config:"envVar"`
	Name       string
	Procedures map[string]string
}

type DatabaseMongo struct {
	Name     string
	Host     string `config:"envVar"`
	User     string `config:"envVar"`
	Password string `config:"envVar"`
}

type Gateways struct {
	Gateway1 Gateway1
}

type Services struct {
	UpdateService UpdateService
}

type UpdateService struct {
	UpdateInterval time.Duration
}
