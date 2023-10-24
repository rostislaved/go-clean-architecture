package apiController

import (
	"log/slog"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/service"
)

type Controller struct {
	Logger  *slog.Logger
	service *service.ApiService
}

func New(logger *slog.Logger, service *service.ApiService) *Controller {
	return &Controller{
		Logger:  logger,
		service: service,
	}
}
