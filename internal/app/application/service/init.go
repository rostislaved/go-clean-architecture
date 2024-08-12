package service

import (
	"context"
	"log/slog"

	"github.com/rostislaved/go-clean-architecture/internal/app/config"
	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

type ApiService struct {
	logger     *slog.Logger
	config     config.UpdateService
	repository repository
	provider   provider
}

type repository interface {
	GetBooks(ctx context.Context, ids []int) (books []book.Book, err error)
	SaveBooks(ctx context.Context, books []book.Book) (ids []int, err error)
}

type provider interface{}

func New(l *slog.Logger, cfg config.UpdateService, repository repository, provider provider) *ApiService {
	return &ApiService{
		logger:     l,
		config:     cfg,
		repository: repository,
		provider:   provider,
	}
}
