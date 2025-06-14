package config

import (
	httpAdapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter"
	kafka_queue "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/kafka-adapter-subscriber/kafka-queue"
	nats_adapter_subscriber "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/nats-adapter-subscriber"
	pprofAdapter "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/pprof-adapter"
	entity5_gateway "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/gateways/entity5-gateway"
	kafka_adapter_publisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/kafka-adapter-publisher"
	nats_adapter_publisher "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/nats-adapter-publisher"
	entity1_repository "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/entity1-repository"
	entity2_repository "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/entity2-repository"
	entity3_repository "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/entity3-repository"
	entity4_repository "github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/repositories/entity4-repository"
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
	Entity1Config entity1_repository.Config
	Entity2Config entity2_repository.Config
	Entity3Config entity3_repository.Config
	Entity4Config entity4_repository.Config

	NatsAdapterPublisher  nats_adapter_publisher.Config
	KafkaAdapterPublisher kafka_adapter_publisher.Config
	Gateways              Gateways
}

type Gateways struct {
	Entity5Gateway entity5_gateway.Config
}
