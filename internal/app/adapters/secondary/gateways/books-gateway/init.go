package books_gateway

import (
	"log/slog"

	"github.com/go-resty/resty/v2"

	providerhelpers "github.com/rostislaved/go-clean-architecture/internal/libs/provider-helpers"
)

type BooksGateway struct {
	logger *slog.Logger
	config Config
	client *resty.Client
}

func New(
	l *slog.Logger,
	config Config,
) *BooksGateway {
	err := providerhelpers.ValidateEndpoints(config.Endpoints)
	if err != nil {
		panic(err)
	}

	client := resty.New().
		SetBaseURL(config.Host).
		// SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetRetryCount(3)

	return &BooksGateway{
		logger: l,
		config: config,
		client: client,
	}
}
