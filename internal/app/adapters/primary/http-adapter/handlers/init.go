package handlers

import (
	"log/slog"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
)

type Handlers struct {
	Logger  *slog.Logger
	service *usecases.UseCases
}

func New(logger *slog.Logger, service *usecases.UseCases) *Handlers {
	return &Handlers{
		Logger:  logger,
		service: service,
	}
}
