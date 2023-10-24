package service

import (
	"github.com/rostislaved/go-clean-architecture/internal/app/domain/config"
)

type ApiService struct {
	logger     *slog.Logger
	config     config.UpdateService
	repository repository
	provider   provider
}

type repository interface{}

type provider interface{}

func New(l *slog.Logger, cfg config.UpdateService, repository repository, provider provider) *ApiService {
	return &ApiService{
		logger:     l,
		config:     cfg,
		repository: repository,
		provider:   provider,
	}
}
