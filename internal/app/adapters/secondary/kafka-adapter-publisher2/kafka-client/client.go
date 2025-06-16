package kafka_client

import (
	"log/slog"
	"strings"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kslog"
)

type Config struct {
	Hosts string `json:"hosts"`
}

func New(l *slog.Logger, cfg Config) (*kgo.Client, error) {
	seeds := strings.Split(cfg.Hosts, ",")

	opts := []kgo.Opt{
		kgo.SeedBrokers(seeds...),
		kgo.WithLogger(kslog.New(l)),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return client, nil
}
