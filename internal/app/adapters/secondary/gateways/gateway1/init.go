package gateway1

import (
	"log"
	"log/slog"

	"github.com/go-resty/resty/v2"

	providerhelpers "github.com/rostislaved/go-clean-architecture/internal/pkg/provider-helpers"
)

type Gateway1 struct {
	logger *slog.Logger
	config Config
	client *resty.Client
}

func New(
	l *slog.Logger,
	config Config,
) *Gateway1 {
	err := providerhelpers.ValidateEndpoints(config.Endpoints)
	if err != nil {
		log.Fatal(err)
	}

	client := resty.New().
		SetBaseURL(config.Host).
		// SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetRetryCount(3)

	return &Gateway1{
		logger: l,
		config: config,
		client: client,
	}
}
