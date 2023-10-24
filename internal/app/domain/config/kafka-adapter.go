package config

type KafkaAdapterSubscriber struct {
	Host    string
	GroupID string
	Topic   string
}
type KafkaAdapterPublisher struct{}
