package usecases

import (
	"context"
	"log/slog"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

type UseCases struct {
	logger                *slog.Logger
	config                Config
	booksRepository       booksRepository
	gateway               gateway
	kafkaAdapterPublisher bookSender
	natsAdapterPublisher  bookSender
}

type Config struct {
	UpdateInterval time.Duration
}

type booksRepository interface {
	Get(ctx context.Context, ids []int) (books []entity1.Entity1, err error)
	Save(ctx context.Context, books []entity1.Entity1) (ids []int, err error)
}

type bookSender interface {
	SendBook(ctx context.Context, b entity1.Entity1) error
}

type gateway interface {
	GetBooks(ctx context.Context, input struct{}) (books []entity1.Entity1, err error)
}

func New(
	l *slog.Logger,
	cfg Config,
	repository booksRepository,
	gateway gateway,
	kafkaAdapterPublisher bookSender,
	natsAdapterPublisher bookSender,
) *UseCases {
	return &UseCases{
		logger:                l,
		config:                cfg,
		booksRepository:       repository,
		gateway:               gateway,
		kafkaAdapterPublisher: kafkaAdapterPublisher,
		natsAdapterPublisher:  natsAdapterPublisher,
	}
}
