package entity5_gateway

import (
	"log/slog"

	"github.com/go-resty/resty/v2"

	providerhelpers "github.com/rostislaved/go-clean-architecture/internal/pkg/provider-helpers"
)

type Entity5Gateway struct {
	logger *slog.Logger
	config Config
	client *resty.Client
}

func New(
	l *slog.Logger,
	config Config,
) *Entity5Gateway {
	err := providerhelpers.ValidateEndpoints(config.Endpoints)
	if err != nil {
		panic(err)
	}

	client := resty.New().
		SetBaseURL(config.Host).
		// SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetRetryCount(3)

	return &Entity5Gateway{
		logger: l,
		config: config,
		client: client,
	}
}
