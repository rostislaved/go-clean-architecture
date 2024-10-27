package kafka_handlers

import (
	"log/slog"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
)

type KafkaHandlers struct {
	Logger  *slog.Logger
	service *usecases.UseCases
}

func New(logger *slog.Logger, service *usecases.UseCases) *KafkaHandlers {
	return &KafkaHandlers{
		Logger:  logger,
		service: service,
	}
}
