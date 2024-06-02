package config

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

type Gateways struct {
	Gateway1 Gateway1
}
