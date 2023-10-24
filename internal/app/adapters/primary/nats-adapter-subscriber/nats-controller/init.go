package natsController

import (
	"log/slog"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/service"
)

type NatsController struct {
	Logger  *slog.Logger
	service *service.ApiService
}

func New(logger *slog.Logger, service *service.ApiService) *NatsController {
	return &NatsController{
		Logger:  logger,
		service: service,
	}
}
