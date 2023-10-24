package kafkaController

import (
	"log/slog"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/service"
)

type KafkaController struct {
	Logger  *slog.Logger
	service *service.ApiService
}

func New(logger *slog.Logger, service *service.ApiService) *KafkaController {
	return &KafkaController{
		Logger:  logger,
		service: service,
	}
}
