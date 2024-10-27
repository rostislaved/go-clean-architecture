package nats_handlers

import (
	"log/slog"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
)

type NatsHandlers struct {
	Logger  *slog.Logger
	service *usecases.UseCases
}

func New(logger *slog.Logger, service *usecases.UseCases) *NatsHandlers {
	return &NatsHandlers{
		Logger:  logger,
		service: service,
	}
}
