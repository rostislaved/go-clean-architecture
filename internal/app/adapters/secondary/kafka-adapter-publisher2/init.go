package kafka_adapter_publisher

import (
	"log/slog"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaAdapterPublisher struct {
	logger *slog.Logger
	config Config
	client *kgo.Client
}
type Config struct{}

func New(logger *slog.Logger, config Config, client *kgo.Client) *KafkaAdapterPublisher {
	return &KafkaAdapterPublisher{
		logger: logger,
		config: config,
		client: client,
	}
}
